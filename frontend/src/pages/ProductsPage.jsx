import { Helmet } from "react-helmet-async";
import { useState, useEffect } from "react";
import axios from "axios";

import {
  Card,
  Table,
  Stack,
  Paper,
  Button,
  Popover,
  Checkbox,
  TableRow,
  MenuItem,
  TableBody,
  TableCell,
  Container,
  Typography,
  IconButton,
  TableContainer,
  TablePagination,
} from "@mui/material";
// components
// import Label from "../components/label";
import Iconify from "../components/iconify/Iconify";
// import Scrollbar from "../components/scrollbar";
// sections
import ProductListHead from "../sections/@dashboard/products/ProductListHead";
import ProductListToolbar from "../sections/@dashboard/products/ProductListToolbar";
// ----------------------------------------------------------------------
import SimpleBarReact from "simplebar-react";
import "simplebar-react/dist/simplebar.min.css";
import AddProductModal from "../components/modal/product/AddProduct";
import EditProduct from "../components/modal/product/Edit";

const TABLE_HEAD = [
  { id: "code", label: "Code", alignRight: false },
  { id: "name", label: "Name", alignRight: false },
  { id: "price", label: "Price", alignRight: false },
  { id: "quantity", label: "Quantity", alignRight: false },
  { id: "type", label: "Type", alignRight: false },
  { id: "" },
];

// ----------------------------------------------------------------------

// function descendingComparator(a, b, orderBy) {
//   if (b[orderBy] < a[orderBy]) {
//     return -1;
//   }
//   if (b[orderBy] > a[orderBy]) {
//     return 1;
//   }
//   return 0;
// }

// function getComparator(order, orderBy) {
//   return order === "desc"
//     ? (a, b) => descendingComparator(a, b, orderBy)
//     : (a, b) => -descendingComparator(a, b, orderBy);
// }

// function applySortFilter(array, comparator, query) {
//   const stabilizedThis = array.map((el, id) => [el, id]);
//   stabilizedThis.sort((a, b) => {
//     const order = comparator(a[0], b[0]);
//     if (order !== 0) return order;
//     return a[1] - b[1];
//   });
//   if (query) {
//     return filter(
//       array,
//       (_user) => _user.name.toLowerCase().indexOf(query.toLowerCase()) !== -1
//     );
//   }
//   return stabilizedThis.map((el) => el[0]);
// }

export default function UserPage() {
  const [open, setOpen] = useState(null);

  const [page, setPage] = useState(0);

  const [order, setOrder] = useState("asc");

  const [selected, setSelected] = useState([]);

  const [orderBy, setOrderBy] = useState("name");

  const [filterName, setFilterName] = useState("");

  const [rowsPerPage, setRowsPerPage] = useState(5);

  const [product, setProduct] = useState([]);

  const [showModal, setShowModal] = useState(false);

  const [editModal, setEditModal] = useState(false);

  const [edit, setEdit] = useState({});

  useEffect(() => {
    console.log("useEffect");
    // console.log(process?.env)
    // console.log(import.meta.env)
    let url = import.meta.env.VITE_MOCK_URL;
    if(import.meta.env.PROD){
      url = 'import.meta.env.VITE_MOCK_URL'
    }
    axios
      .get(url + "products")
      .then((response) => {
        setProduct(response.data);
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  const handleOpenMenu = (event) => {
    setOpen(event.currentTarget);
  };

  const handleCloseMenu = () => {
    setOpen(null);
  };

  const handleRequestSort = (event, property) => {
    const isAsc = orderBy === property && order === "asc";
    setOrder(isAsc ? "desc" : "asc");
    setOrderBy(property);
  };

  const handleSelectAllClick = (event) => {
    if (event.target.checked) {
      const newSelecteds = product.map((n) => n.id);
      setSelected(newSelecteds);
      return;
    }
    setSelected([]);
  };

  const handleClick = (event, name) => {
    const selectedIndex = selected.indexOf(name);
    let newSelected = [];
    if (selectedIndex === -1) {
      newSelected = newSelected.concat(selected, name);
    } else if (selectedIndex === 0) {
      newSelected = newSelected.concat(selected.slice(1));
    } else if (selectedIndex === selected.length - 1) {
      newSelected = newSelected.concat(selected.slice(0, -1));
    } else if (selectedIndex > 0) {
      newSelected = newSelected.concat(
        selected.slice(0, selectedIndex),
        selected.slice(selectedIndex + 1)
      );
    }
    setSelected(newSelected);
  };

  const handleChangePage = (event, newPage) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (event) => {
    setPage(0);
    setRowsPerPage(parseInt(event.target.value, 10));
  };

  const handleFilterByName = (event) => {
    setPage(0);
    setFilterName(event.target.value);
  };

  const handleClickOpen = () => {
    setShowModal(!showModal);
  };

  const handleEdit = () => {
    setEditModal(!editModal);
    setOpen(null);
  };

  const handleDelete = () => {
    setEditModal(!editModal);
    setOpen(null);
  };

  const emptyRows =
    page > 0 ? Math.max(0, (1 + page) * rowsPerPage - product.length) : 0;

  // const filteredProducts = applySortFilter(
  //   product,
  //   getComparator(order, orderBy),
  //   filterName
  // );

  const isNotFound = !product.length && !!filterName;

  return (
    <>
      <Helmet>
        <title> Product | InventoryHub </title>
      </Helmet>

      <Container>
        <Stack
          direction="row"
          alignItems="center"
          justifyContent="space-between"
          mb={5}
        >
          <Typography variant="h4" gutterBottom>
            Product List
          </Typography>
          {product.length != 0 && (
            <Button
              variant="contained"
              startIcon={<Iconify icon="eva:plus-fill" />}
              onClick={handleClickOpen}
            >
              Add Product
            </Button>
          )}
        </Stack>

        <Card>
          <ProductListToolbar
            numSelected={selected.length}
            filterName={filterName}
            onFilterName={handleFilterByName}
            product={selected}
          />

          <TableContainer sx={{ minWidth: 800, height: 450 }}>
            <SimpleBarReact style={{ maxHeight: 450 }}>
              <Table>
                <ProductListHead
                  order={order}
                  orderBy={orderBy}
                  headLabel={TABLE_HEAD}
                  rowCount={product.length}
                  numSelected={selected.length}
                  onRequestSort={handleRequestSort}
                  onSelectAllClick={handleSelectAllClick}
                />

                <TableBody>
                  {product
                    ?.slice(
                      page * rowsPerPage,
                      page * rowsPerPage + rowsPerPage
                    )
                    ?.map((row) => {
                      const { id, code, name, price, type, quantity } = row;
                      const selectedProduct = selected.indexOf(id) !== -1;

                      return (
                        <TableRow
                          hover
                          key={id}
                          tabIndex={-1}
                          role="checkbox"
                          selected={selectedProduct}
                        >
                          <TableCell padding="checkbox">
                            <Checkbox
                              checked={selectedProduct}
                              onChange={(event) => handleClick(event, id)}
                            />
                          </TableCell>
                          <TableCell align="left">{code}</TableCell>
                          <TableCell align="left">{name}</TableCell>
                          <TableCell align="left">{price}</TableCell>
                          <TableCell align="left">{quantity}</TableCell>
                          <TableCell align="left">{type}</TableCell>

                          <TableCell align="right">
                            <IconButton
                              size="large"
                              color="inherit"
                              onClick={(event) => {
                                handleOpenMenu(event);
                                setEdit(row);
                              }}
                            >
                              <Iconify icon={"eva:more-vertical-fill"} />
                            </IconButton>
                          </TableCell>
                        </TableRow>
                      );
                    })}
                  {emptyRows > 0 && (
                    <TableRow style={{ height: 53 * emptyRows }}>
                      <TableCell colSpan={6} />
                    </TableRow>
                  )}
                </TableBody>

                {isNotFound && (
                  <TableBody>
                    <TableRow>
                      <TableCell align="center" colSpan={6} sx={{ py: 3 }}>
                        <Paper
                          sx={{
                            textAlign: "center",
                          }}
                        >
                          <Typography variant="h6" paragraph>
                            Not found
                          </Typography>

                          <Typography variant="body2">
                            No results found for &nbsp;
                            <strong>&quot;{filterName}&quot;</strong>.
                            <br /> Try checking for typos or using complete
                            words.
                          </Typography>
                        </Paper>
                      </TableCell>
                    </TableRow>
                  </TableBody>
                )}

                {product.length == 0 && (
                  <TableBody>
                    <TableRow>
                      <TableCell align="center" colSpan={6} sx={{ py: 3 }}>
                        <Paper
                          sx={{
                            textAlign: "center",
                          }}
                        >
                          <Typography variant="h6" paragraph>
                            No data, create the first one!
                          </Typography>
                          <Button
                            variant="contained"
                            startIcon={<Iconify icon="eva:plus-fill" />}
                            onClick={handleClickOpen}
                          >
                            Add Product
                          </Button>
                        </Paper>
                      </TableCell>
                    </TableRow>
                  </TableBody>
                )}
              </Table>
            </SimpleBarReact>
          </TableContainer>

          <TablePagination
            rowsPerPageOptions={[5, 10, 25]}
            component="div"
            count={product.length}
            rowsPerPage={rowsPerPage}
            page={page}
            onPageChange={handleChangePage}
            onRowsPerPageChange={handleChangeRowsPerPage}
          />
        </Card>
      </Container>

      {/* Modal */}
      <AddProductModal isOpen={showModal} onClose={() => setShowModal(false)} />

      <EditProduct
        isOpen={editModal}
        onClose={() => setEditModal(false)}
        data={edit}
      />

      <Popover
        open={Boolean(open)}
        anchorEl={open}
        onClose={handleCloseMenu}
        anchorOrigin={{
          vertical: "top",
          horizontal: "left",
        }}
        transformOrigin={{
          vertical: "top",
          horizontal: "right",
        }}
        PaperProps={{
          sx: {
            p: 1,
            width: 140,
            "& .MuiMenuItem-root": {
              px: 1,
              typography: "body2",
              borderRadius: 0.75,
            },
          },
        }}
      >
        <MenuItem onClick={handleEdit}>
          <Iconify icon={"eva:edit-fill"} sx={{ mr: 2 }} />
          Edit
        </MenuItem>

        <MenuItem sx={{ color: "error.main" }} onClick={handleDelete}>
          <Iconify icon={"eva:trash-2-outline"} sx={{ mr: 2 }} />
          Delete
        </MenuItem>
      </Popover>
    </>
  );
}
