import { useState } from "react";
import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
  TextField,
  Grid,
  MenuItem,
  Box,
} from "@mui/material";
import PropTypes from "prop-types";

import "./AddProduct.css";

const AddProductModal = ({ isOpen, onClose }) => {
  const [code, setCode] = useState("");
  const [name, setName] = useState("");
  const [currency, setCurrency] = useState("IDR");
  const [price, setPrice] = useState("");
  const [type, setType] = useState("");
  const [quantity, setQuantity] = useState("");
  const [unit, setUnit] = useState("pcs");

  const currencies = [
    {
      value: "IDR",
      label: "Rp",
    },
    {
      value: "USD",
      label: "$",
    },
    {
      value: "EUR",
      label: "€",
    },
  ];

  const types = [
    {
      value: "",
      label: "None",
    },
    {
      value: "automotive",
      label: "Automotive",
    },
    {
      value: "electronic",
      label: "Electronic",
    },
    {
      value: "fashion",
      label: "Fashion",
    },
    {
      value: "kids",
      label: "Kids",
    },
  ];

  const units = [
    {
      value: "pcs",
      label: "pcs",
    },
    {
      value: "kg",
      label: "kg",
    },
    {
      value: "g",
      label: "g",
    },
  ];

  const handleAddProduct = () => {
    const newProduct = {
      code: code,
      name: name,
      currency: currency,
      price: parseInt(price),
      type: type,
      quantity: parseInt(quantity),
      unit: unit,
    };

    console.log(newProduct);

    onSubmit(newProduct);

    // Reset the form inputs
    setCode("");
    setName("");
    setPrice("");
    setType("");
    setQuantity("");
  };

  const onSubmit = async (data) => {
    try {
      const response = await fetch(import.meta.env.VITE_API_URL + "/products", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });

      if (response.ok) {
        alert("Successfully add new product");
        // setProduct([...product, data]);
      } else {
        // Handle error response here if needed
        alert("Error sending product data to the API");
      }
      onClose(true);
    } catch (error) {
      // Handle any network errors or other issues
      alert("Error sending product data:", error);
    }
  };

  if (!isOpen) {
    return null;
  }

  const handleClose = () => {
    //setOpen(false);
    onClose(false);
  };

  const validation = (text) => text.length < 1;

  const isValid =
    validation(code) ||
    validation(name) ||
    validation(price) ||
    validation(type) ||
    validation(quantity);

  console.log({ isValid });

  return (
    <>
      <Dialog
        open={isOpen}
        onClose={handleClose}
        fullWidth={true}
        // maxWidth={maxWidth}
      >
        <DialogTitle>Add Product</DialogTitle>
        <DialogContent>
          <Box
            noValidate
            component="form"
            sx={{
              display: "flex",
              flexDirection: "column",
              m: "auto",
              width: "fit-content",
            }}
          >
            <Grid
              container
              rowSpacing={1}
              columnSpacing={{ xs: 1, sm: 2, md: 3 }}
            >
              <Grid item xs={12} md={6}>
                <TextField
                  autoFocus
                  margin="dense"
                  id="code"
                  label="Code"
                  type="text"
                  fullWidth
                  variant="outlined"
                  onChange={(e) => {
                    setCode(e.target.value);
                  }}
                />
              </Grid>
              <Grid item xs={12} md={6}>
                <TextField
                  autoFocus
                  margin="dense"
                  id="name"
                  label="Name"
                  type="text"
                  fullWidth
                  variant="outlined"
                  onChange={(e) => {
                    setName(e.target.value);
                  }}
                />
              </Grid>
              <Grid item xs={3} md={2}>
                <TextField
                  id="currency"
                  select
                  value={currency}
                  margin="dense"
                  fullWidth
                  onChange={(e) => {
                    setCurrency(e.target.value);
                  }}
                >
                  {currencies.map((option) => (
                    <MenuItem key={option.value} value={option.value}>
                      {option.label}
                    </MenuItem>
                  ))}
                </TextField>
              </Grid>
              <Grid item xs={9} md={10}>
                <TextField
                  autoFocus
                  margin="dense"
                  id="price"
                  label="Price"
                  type="number"
                  fullWidth
                  variant="outlined"
                  onChange={(e) => {
                    setPrice(e.target.value);
                  }}
                />
              </Grid>
              <Grid item xs={12} md={5}>
                <TextField
                  id="type"
                  select
                  defaultValue=""
                  margin="dense"
                  label="Type"
                  fullWidth
                  onChange={(e) => {
                    setType(e.target.value);
                  }}
                >
                  {types.map((option) => (
                    <MenuItem key={option.value} value={option.value}>
                      {option.label}
                    </MenuItem>
                  ))}
                </TextField>
              </Grid>
              <Grid item xs={8} md={4}>
                <TextField
                  autoFocus
                  margin="dense"
                  id="quantity"
                  label="Quantity"
                  type="number"
                  fullWidth
                  variant="outlined"
                  onChange={(e) => {
                    setQuantity(e.target.value);
                  }}
                />
              </Grid>
              <Grid item xs={4} md={3}>
                <TextField
                  id="unit"
                  select
                  value={unit}
                  margin="dense"
                  fullWidth
                  onChange={(e) => {
                    setUnit(e.target.value);
                  }}
                >
                  {units.map((option) => (
                    <MenuItem key={option.value} value={option.value}>
                      {option.label}
                    </MenuItem>
                  ))}
                </TextField>
              </Grid>
            </Grid>
          </Box>
          <DialogActions>
            <Button variant="outlined" onClick={handleClose}>
              Cancel
            </Button>
            <Button
              variant="contained"
              onClick={handleAddProduct}
              disabled={isValid}
            >
              Submit
            </Button>
          </DialogActions>
        </DialogContent>
      </Dialog>
    </>
  );
};

AddProductModal.propTypes = {
  isOpen: PropTypes.bool,
  onClose: PropTypes.func,
};

export default AddProductModal;
