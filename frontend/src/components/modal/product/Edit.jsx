import { useEffect, useState } from "react";
import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
  // TextField,
} from "@mui/material";
import PropTypes from "prop-types";

// import "./Edit.css";

const onSubmit = async (data) => {
  console.log(data)
  try {
    // Make the API request to send the product data
    const response = await fetch(
      import.meta.env.VITE_API_URL+"/products/edit/" + data?.id,
      {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      }
    );

    if (response.ok) {
      // If the API call is successful, add the product to the state
      alert("Successfully add new product");
      // setProduct([...product, data]);
    } else {
      // Handle error response here if needed
      alert("Error sending product data to the API");
    }
    // setShowModal(false);
  } catch (error) {
    // Handle any network errors or other issues
    alert("Error sending product data:", error);
  }

  // Close the modal after submitting the form
  // setProducts([...products, product]);
};

const EditProduct = ({ isOpen, onClose, data }) => {
  // const [code, setCode] = useState("");
  // const [name, setName] = useState("");
  // const [price, setPrice] = useState("");
  // const [type, setType] = useState("");
  // const [quantity, setQuantity] = useState("");

  const [fullWidth] = useState(true);
  const [maxWidth] = useState("sm");

  const [editedData, setEditedData] = useState(data);

  useEffect(() => {
    setEditedData(data);
  }, [data]);

  // console.log(data)
  // setCode(data?.code)

  // const handleCodeChange = (e) => {
  //   setCode(e.target.value);
  // };

  // const handleProductNameChange = (e) => {
  //   setName(e.target.value);
  // };

  // const handlePriceChange = (e) => {
  //   setPrice(e.target.value);
  // };

  // const handleQuantityChange = (e) => {
  //   setQuantity(e.target.value);
  // };

  // const handleTypeChange = (e) => {
  //   setType(e.target.value);
  // };
// const setData = ()=>{
//   setEditedData(data)
// }

  const handleInput = (e) => {
    const { name, value } = e.target;
    setEditedData((data) => ({
      ...data,
      [name]: value,
    }));
    // console.log({ code, name, price, type, quantity });
    // setEditedData({ code, name, price, type, quantity })
  };

  const handleAddProduct = (e) => {
    e.preventDefault();

    // if (!code || !name || !price || !type || !quantity) {
    //   alert("Please filled all required field");
    //   return;
    // }

    onSubmit(editedData);

    // Reset the form inputs
    // setCode("");
    // setName("");
    // setPrice("");
    // setType("");
    // setQuantity("");

    onClose(true);
  };

  if (!isOpen) {
    return null;
  }

  const handleClose = () => {
    //setOpen(false);
    onClose(false);
  };

  return (
    <>
      <Dialog
        open={isOpen}
        onClose={handleClose}
        fullWidth={fullWidth}
        maxWidth={maxWidth}
      >
        <DialogTitle>Edit Product</DialogTitle>
        <DialogContent>
          <div className="field">
            <label className="label">Product Code</label>
            <div className="control">
              <input
                className="input"
                name="code"
                type="text"
                defaultValue={editedData.code ? editedData.code : data.code}
                value={editedData.code}
                onChange={handleInput}
              />
            </div>
          </div>

          <div className="field">
            <label className="label">Product Name</label>
            <div className="control">
              <input
                className="input"
                name="name"
                type="text"
                defaultValue={data?.name}
                value={editedData?.name}
                onChange={handleInput}
              />
            </div>
          </div>

          <div className="field">
            <label className="label">Price</label>
            <div className="control">
              <input
                className="input"
                name="price"
                type="number"
                defaultValue={data?.price || ""}
                value={editedData?.price}
                onChange={handleInput}
              />
            </div>
          </div>

          <div className="field">
            <label className="label">Quantity</label>
            <div className="control">
              <input
                className="input"
                name="quantity"
                type="number"
                defaultValue={data?.quantity || ""}
                value={editedData?.quantity}
                onChange={handleInput}
              />
            </div>
          </div>

          <div className="field">
            <label className="label">Type</label>
            <div className="control">
              <input
                className="input"
                name="type"
                type="text"
                defaultValue={data?.type || ""}
                value={editedData?.type}
                onChange={handleInput}
              />
            </div>
          </div>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose}>Cancel</Button>
          <Button onClick={handleAddProduct}>Edit</Button>
        </DialogActions>
      </Dialog>
    </>
  );
};

EditProduct.propTypes = {
  isOpen: PropTypes.bool,
  onClose: PropTypes.func,
  data: PropTypes.object,
};

export default EditProduct;
