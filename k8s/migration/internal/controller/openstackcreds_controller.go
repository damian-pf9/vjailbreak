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

	"github.com/go-logr/logr"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumetypes"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	vjailbreakv1alpha1 "github.com/platform9/vjailbreak/k8s/migration/api/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// OpenstackCredsReconciler reconciles a OpenstackCreds object
type OpenstackCredsReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

type OpenStackClients struct {
	BlockStorageClient *gophercloud.ServiceClient
	ComputeClient      *gophercloud.ServiceClient
	NetworkingClient   *gophercloud.ServiceClient
}

// +kubebuilder:rbac:groups=vjailbreak.k8s.pf9.io,resources=openstackcreds,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=vjailbreak.k8s.pf9.io,resources=openstackcreds/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=vjailbreak.k8s.pf9.io,resources=openstackcreds/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the OpenstackCreds object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.18.4/pkg/reconcile
func (r *OpenstackCredsReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	ctxlog := log.FromContext(ctx)

	// Get the OpenstackCreds object
	openstackcreds := &vjailbreakv1alpha1.OpenstackCreds{}
	if err := r.Get(ctx, req.NamespacedName, openstackcreds); err != nil {
		if apierrors.IsNotFound(err) {
			ctxlog.Info("Received ignorable event for a recently deleted OpenstackCreds.")
			return ctrl.Result{}, nil
		}
		ctxlog.Error(err, fmt.Sprintf("Unexpected error reading OpenstackCreds '%s' object", openstackcreds.Name))
		return ctrl.Result{}, err
	}

	if openstackcreds.ObjectMeta.DeletionTimestamp.IsZero() {
		// Check if speck matches with kubectl.kubernetes.io/last-applied-configuration
		ctxlog.Info(fmt.Sprintf("OpenstackCreds '%s' CR is being created or updated", openstackcreds.Name))
		ctxlog.Info(fmt.Sprintf("Validating OpenstackCreds '%s' object", openstackcreds.Name))
		if _, err := validateOpenstackCreds(ctxlog, openstackcreds); err != nil {
			// Update the status of the OpenstackCreds object
			openstackcreds.Status.OpenStackValidationStatus = "Failed"
			openstackcreds.Status.OpenStackValidationMessage = fmt.Sprintf("Error validating OpenstackCreds '%s'", openstackcreds.Name)
			if err := r.Status().Update(ctx, openstackcreds); err != nil {
				ctxlog.Error(err, fmt.Sprintf("Error updating status of OpenstackCreds '%s'", openstackcreds.Name))
				return ctrl.Result{}, err
			}
		} else {
			ctxlog.Info(fmt.Sprintf("Successfully authenticated to Openstack '%s'", openstackcreds.Spec.OsAuthURL))
			// Update the status of the OpenstackCreds object
			openstackcreds.Status.OpenStackValidationStatus = "Succeeded"
			openstackcreds.Status.OpenStackValidationMessage = "Successfully authenticated to Openstack"
			if err := r.Status().Update(ctx, openstackcreds); err != nil {
				ctxlog.Error(err, fmt.Sprintf("Error updating status of OpenstackCreds '%s'", openstackcreds.Name))
				return ctrl.Result{}, err
			}
		}
	}
	return ctrl.Result{}, nil
}

func validateOpenstackCreds(ctxlog logr.Logger, openstackcreds *vjailbreakv1alpha1.OpenstackCreds) (*OpenStackClients, error) {
	providerClient, err := openstack.AuthenticatedClient(gophercloud.AuthOptions{
		IdentityEndpoint: openstackcreds.Spec.OsAuthURL,
		Username:         openstackcreds.Spec.OsUsername,
		Password:         openstackcreds.Spec.OsPassword,
		DomainName:       openstackcreds.Spec.OsDomainName,
		TenantName:       openstackcreds.Spec.OsTenantName,
	})
	if err != nil {
		ctxlog.Error(err, fmt.Sprintf("Error authenticating to Openstack '%s'", openstackcreds.Spec.OsAuthURL))
		return nil, err
	}
	endpoint := gophercloud.EndpointOpts{
		Region: openstackcreds.Spec.OsRegionName,
	}
	computeClient, err := openstack.NewComputeV2(providerClient, endpoint)
	if err != nil {
		ctxlog.Error(err, fmt.Sprintf("Error validating region '%s' for '%s'",
			openstackcreds.Spec.OsRegionName, openstackcreds.Spec.OsAuthURL))
		return nil, err
	}
	blockStorageClient, err := openstack.NewBlockStorageV3(providerClient, endpoint)
	if err != nil {
		ctxlog.Error(err, fmt.Sprintf("Error validating region '%s' for '%s'",
			openstackcreds.Spec.OsRegionName, openstackcreds.Spec.OsAuthURL))
		return nil, err
	}
	networkingClient, err := openstack.NewNetworkV2(providerClient, endpoint)
	if err != nil {
		ctxlog.Error(err, fmt.Sprintf("Error validating region '%s' for '%s'",
			openstackcreds.Spec.OsRegionName, openstackcreds.Spec.OsAuthURL))
		return nil, err
	}
	return &OpenStackClients{
		BlockStorageClient: blockStorageClient,
		ComputeClient:      computeClient,
		NetworkingClient:   networkingClient,
	}, nil
}

//nolint:dupl // This function is similar to VerifyNetworks, excluding from linting to keep it readable
func VerifyNetworks(ctx context.Context, openstackcreds *vjailbreakv1alpha1.OpenstackCreds, targetnetworks []string) error {
	openstackClients, err := validateOpenstackCreds(log.FromContext(ctx), openstackcreds)
	if err != nil {
		return err
	}
	allPages, err := networks.List(openstackClients.NetworkingClient, nil).AllPages()
	if err != nil {
		return fmt.Errorf("failed to list networks: %w", err)
	}

	allNetworks, err := networks.ExtractNetworks(allPages)
	if err != nil {
		return fmt.Errorf("failed to extract all networks: %w", err)
	}

	// Verify that all network names in targetnetworks exist in the openstack networks
	for _, targetNetwork := range targetnetworks {
		found := false
		for i := 0; i < len(allNetworks); i++ {
			if allNetworks[i].Name == targetNetwork {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("network '%s' not found in OpenStack", targetNetwork)
		}
	}
	return nil
}

//nolint:dupl // This function is similar to VerifyNetworks, excluding from linting to keep it readable
func VerifyStorage(ctx context.Context, openstackcreds *vjailbreakv1alpha1.OpenstackCreds, targetstorages []string) error {
	openstackClients, err := validateOpenstackCreds(log.FromContext(ctx), openstackcreds)
	if err != nil {
		return err
	}
	allPages, err := volumetypes.List(openstackClients.BlockStorageClient, nil).AllPages()
	if err != nil {
		return fmt.Errorf("failed to list volume types: %w", err)
	}

	allvoltypes, err := volumetypes.ExtractVolumeTypes(allPages)
	if err != nil {
		return fmt.Errorf("failed to extract all volume types: %w", err)
	}

	// Verify that all volume types in targetstorage exist in the openstack volume types
	for _, targetstorage := range targetstorages {
		found := false
		for i := 0; i < len(allvoltypes); i++ {
			if allvoltypes[i].Name == targetstorage {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("volume type '%s' not found in OpenStack", targetstorage)
		}
	}
	return nil
}

func GetOpenstackInfo(ctx context.Context, openstackcreds *vjailbreakv1alpha1.OpenstackCreds) (*vjailbreakv1alpha1.OpenstackInfo, error) {
	var openstackvoltypes []string
	var openstacknetworks []string
	openstackClients, err := validateOpenstackCreds(log.FromContext(ctx), openstackcreds)
	if err != nil {
		return nil, err
	}
	allVolumeTypePages, err := volumetypes.List(openstackClients.BlockStorageClient, nil).AllPages()
	if err != nil {
		return nil, fmt.Errorf("failed to list volume types: %w", err)
	}

	allvoltypes, err := volumetypes.ExtractVolumeTypes(allVolumeTypePages)
	if err != nil {
		return nil, fmt.Errorf("failed to extract all volume types: %w", err)
	}

	for i := 0; i < len(allvoltypes); i++ {
		openstackvoltypes = append(openstackvoltypes, allvoltypes[i].Name)
	}

	allNetworkPages, err := networks.List(openstackClients.NetworkingClient, nil).AllPages()
	if err != nil {
		return nil, fmt.Errorf("failed to list networks: %w", err)
	}

	allNetworks, err := networks.ExtractNetworks(allNetworkPages)
	if err != nil {
		return nil, fmt.Errorf("failed to extract all networks: %w", err)
	}

	for i := 0; i < len(allNetworks); i++ {
		openstacknetworks = append(openstacknetworks, allNetworks[i].Name)
	}

	return &vjailbreakv1alpha1.OpenstackInfo{
		VolumeTypes: openstackvoltypes,
		Networks:    openstacknetworks,
	}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *OpenstackCredsReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&vjailbreakv1alpha1.OpenstackCreds{}).
		WithEventFilter(predicate.GenerationChangedPredicate{}).
		Complete(r)
}
