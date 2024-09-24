/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	vjailbreakv1alpha1 "github.com/platform9/vjailbreak/k8s/migration/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// MigrationTemplateReconciler reconciles a MigrationTemplate object
type MigrationTemplateReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	ctxlog logr.Logger
}

// +kubebuilder:rbac:groups=vjailbreak.k8s.pf9.io,resources=migrationtemplates,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=vjailbreak.k8s.pf9.io,resources=migrationtemplates/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=vjailbreak.k8s.pf9.io,resources=migrationtemplates/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the MigrationTemplate object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.18.4/pkg/reconcile
func (r *MigrationTemplateReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.ctxlog = log.FromContext(ctx)
	r.ctxlog.Info("Reconciling MigrationTemplate")

	migrationtemplate := &vjailbreakv1alpha1.MigrationTemplate{}

	if err := r.Get(ctx, req.NamespacedName, migrationtemplate); err != nil {
		if apierrors.IsNotFound(err) {
			r.ctxlog.Info("Received ignorable event for a recently deleted MigrationTemplate.")
			return ctrl.Result{}, nil
		}
		r.ctxlog.Error(err, fmt.Sprintf("Unexpected error reading MigrationTemplate '%s' object", migrationtemplate.Name))
		return ctrl.Result{}, err
	}

	vmwcreds := &vjailbreakv1alpha1.VMwareCreds{}
	if ok, err := r.checkStatusSuccess(ctx, migrationtemplate.Namespace, migrationtemplate.Spec.Source.VMwareRef, true, vmwcreds); !ok {
		return ctrl.Result{
			RequeueAfter: time.Minute,
		}, err
	}
	// Fetch OpenStackCreds CR
	openstackcreds := &vjailbreakv1alpha1.OpenstackCreds{}
	if ok, err := r.checkStatusSuccess(ctx, migrationtemplate.Namespace, migrationtemplate.Spec.Destination.OpenstackRef,
		false, openstackcreds); !ok {
		return ctrl.Result{
			RequeueAfter: time.Minute,
		}, err
	}

	vminfo, err := GetAllVMs(ctx, vmwcreds, migrationtemplate.Spec.Source.DataCenter)
	if err != nil {
		r.ctxlog.Error(err, "Failed to get info of all VMs")
		return ctrl.Result{}, err
	}

	openstackinfo, err := GetOpenstackInfo(ctx, openstackcreds)
	if err != nil {
		r.ctxlog.Error(err, "Failed to get info of Openstack")
		return ctrl.Result{}, err
	}

	// Update the status of the MigrationTemplate CR
	migrationtemplate.Status.VMWare = vminfo
	migrationtemplate.Status.Openstack = *openstackinfo
	if err := r.Status().Update(ctx, migrationtemplate); err != nil {
		r.ctxlog.Error(err, "Failed to update MigrationTemplate status")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *MigrationTemplateReconciler) checkStatusSuccess(ctx context.Context,
	namespace, credsname string,
	isvmware bool,
	credsobj client.Object) (bool, error) {
	err := r.Get(ctx, types.NamespacedName{Name: credsname, Namespace: namespace}, credsobj)
	if err != nil {
		return false, fmt.Errorf("failed to get VMwareCreds: %w", err)
	}

	if isvmware && credsobj.(*vjailbreakv1alpha1.VMwareCreds).Status.VMwareValidationStatus != string(corev1.PodSucceeded) {
		return false, fmt.Errorf("VMwareCreds '%s' CR is not validated", credsobj.(*vjailbreakv1alpha1.VMwareCreds).Name)
	} else if !isvmware && credsobj.(*vjailbreakv1alpha1.OpenstackCreds).Status.OpenStackValidationStatus != string(corev1.PodSucceeded) {
		return false, fmt.Errorf("OpenstackCreds '%s' CR is not validated", credsobj.(*vjailbreakv1alpha1.OpenstackCreds).Name)
	}
	return true, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MigrationTemplateReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		// Uncomment the following line adding a pointer to an instance of the controlled resource as an argument
		For(&vjailbreakv1alpha1.MigrationTemplate{}).
		Complete(r)
}
