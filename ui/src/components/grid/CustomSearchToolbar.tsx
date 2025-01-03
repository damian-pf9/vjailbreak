import { Box, IconButton, Tooltip, Typography } from "@mui/material"
import { GridToolbarQuickFilter } from "@mui/x-data-grid"
import { RefreshRounded } from "@mui/icons-material"

interface CustomSearchToolbarProps {
  title?: string;
  onRefresh?: () => void;
  disableRefresh?: boolean;
}

const CustomSearchToolbar = ({
  title,
  onRefresh,
  disableRefresh = false
}: CustomSearchToolbarProps) => {
  return (
    <Box
      sx={{
        p: 1,
        display: "flex",
        alignItems: "center",
        marginLeft: 2,
        marginRight: 2,
      }}
    >
      {title && <Typography variant="h6">{title}</Typography>}
      <Box sx={{ marginLeft: "auto", display: "flex", gap: 1, alignItems: "center" }}>
        {onRefresh && (
          <Tooltip title="Refresh">
            <span>
              <IconButton
                onClick={onRefresh}
                disabled={disableRefresh}
                size="small"
              >
                <RefreshRounded />
              </IconButton>
            </span>
          </Tooltip>
        )}
        <GridToolbarQuickFilter />
      </Box>
    </Box>
  )
}

export default CustomSearchToolbar
