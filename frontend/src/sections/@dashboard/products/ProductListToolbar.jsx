import PropTypes from "prop-types";
// @mui
import { styled, alpha } from "@mui/material/styles";
import {
  Toolbar,
  Tooltip,
  IconButton,
  Typography,
  OutlinedInput,
  InputAdornment,
} from "@mui/material";
// component
import Iconify from "../../../components/iconify/Iconify";

// ----------------------------------------------------------------------

const StyledRoot = styled(Toolbar)(({ theme }) => ({
  height: 96,
  display: "flex",
  justifyContent: "space-between",
  padding: theme.spacing(0, 1, 0, 3),
}));

const StyledSearch = styled(OutlinedInput)(({ theme }) => ({
  width: 240,
  transition: theme.transitions.create(["box-shadow", "width"], {
    easing: theme.transitions.easing.easeInOut,
    duration: theme.transitions.duration.shorter,
  }),
  "&.Mui-focused": {
    width: 320,
    boxShadow: theme.customShadows.z8,
  },
  "& fieldset": {
    borderWidth: `1px !important`,
    borderColor: `${alpha(theme.palette.grey[500], 0.32)} !important`,
  },
}));

const handleProductDelete = async (product) => {
  // try {
  //   // Make the API request to send the product data
  //   const response = await fetch(import.meta.env.VITE_API_URL+ "/products", {
  //     method: "POST",
  //     headers: {
  //       "Content-Type": "application/json",
  //     },
  //     body: JSON.stringify(product),
  //   });

  //   if (response.ok) {
  //     // If the API call is successful, add the product to the state
  //     alert("Successfully add new product");
  //   } else {
  //     // Handle error response here if needed
  //     alert("Error sending product data to the API");
  //   }
  // } catch (error) {
  //   // Handle any network errors or other issues
  //   alert("Error sending product data:", error);
  // }

  // Close the modal after submitting the form
  // setProducts([...products, product]);
  console.log(product);
};

// ----------------------------------------------------------------------

UserListToolbar.propTypes = {
  numSelected: PropTypes.number,
  filterName: PropTypes.string,
  onFilterName: PropTypes.func,
};

export default function UserListToolbar({
  numSelected,
  filterName,
  onFilterName,
}) {
  return (
    <StyledRoot
      sx={{
        ...(numSelected > 0 && {
          color: "primary.main",
          bgcolor: "primary.lighter",
        }),
      }}
    >
      {numSelected > 0 ? (
        <Typography component="div" variant="subtitle1">
          {numSelected} selected
        </Typography>
      ) : (
        <StyledSearch
          value={filterName}
          onChange={onFilterName}
          placeholder="Search"
          startAdornment={
            <InputAdornment position="start">
              <Iconify
                icon="eva:search-fill"
                sx={{ color: "text.disabled", width: 20, height: 20 }}
              />
            </InputAdornment>
          }
        />
      )}

      {numSelected > 0 ? (
        <Tooltip title="Delete">
          <IconButton
            size="large"
            color="inherit"
            onClick={handleProductDelete}
          >
            <Iconify icon="eva:trash-2-fill" />
          </IconButton>
        </Tooltip>
      ) : null}
    </StyledRoot>
  );
}
