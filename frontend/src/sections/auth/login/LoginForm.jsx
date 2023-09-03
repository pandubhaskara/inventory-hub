import { useState } from "react";
import "isomorphic-fetch";
import { useNavigate } from "react-router-dom";
// @mui
import { Stack, IconButton, InputAdornment, TextField } from "@mui/material";
import { LoadingButton } from "@mui/lab";
// components
import Iconify from "../../../components/iconify/Iconify";
import ErrorModal from "../../../components/modal/error/error";
// ----------------------------------------------------------------------
import axiosPublic from "../../../api/api";
// import useAuth from "../../../hooks/useAuth";

export default function LoginForm() {
  // const { setAuth } = useAuth();

  const navigate = useNavigate();

  const [showPassword, setShowPassword] = useState(false);
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  // const [token, setToken] = useState();
  // const [refreshToken, setRefreshToken] = useState();

  const [errorModalOpen, setErrorModalOpen] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");

  const handleUsernameChange = (event) => {
    setUsername(event.target.value);
  };

  const handlePasswordChange = (event) => {
    setPassword(event.target.value);
  };

  const handleCloseErrorModal = () => {
    setErrorModalOpen(false);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    // Perform login logic with username and password values
    if (!username) {
      setErrorMessage("Username is required");
      setErrorModalOpen(true);
    } else if (!password) {
      setErrorMessage("Password is required");
      setErrorModalOpen(true);
    } else {
      const response = await axiosPublic
        .post("/api/auth/login", {
          username: username,
          password: password,
        })
        .then((response) => {
          return response;
        })
        .catch((err) => {
          if (!err?.response?.data?.message) {
            setErrorMessage("No Server Response");
            setErrorModalOpen(true);
          } else if (err.response?.status === 400) {
            setErrorMessage("Missing Username or Password");
            setErrorModalOpen(true);
          } else if (err.response?.status === 401) {
            setErrorMessage("Unauthorized");
            setErrorModalOpen(true);
          } else {
            setErrorMessage(err?.response?.data?.message);
            setErrorModalOpen(true);
          }
        });

      console.log(JSON.stringify(response?.data));

      // const accessToken = response?.data?.accessToken;
      // const isAdmin = response?.data?.isAdmin;

      // setAuth({ username, password, isAdmin, accessToken });
      setUsername("");
      setPassword("");

      navigate("/dashboard/products");
    }
  };

  return (
    <>
      <Stack spacing={3}>
        <TextField
          name="username"
          label="Username"
          value={username}
          onChange={handleUsernameChange}
        />

        <TextField
          name="password"
          label="Password"
          type={showPassword ? "text" : "password"}
          value={password}
          onChange={handlePasswordChange}
          InputProps={{
            endAdornment: (
              <InputAdornment position="end">
                <IconButton
                  onClick={() => setShowPassword(!showPassword)}
                  edge="end"
                >
                  <Iconify
                    icon={showPassword ? "eva:eye-fill" : "eva:eye-off-fill"}
                  />
                </IconButton>
              </InputAdornment>
            ),
          }}
        />
      </Stack>
      <Stack
        direction="row"
        alignItems="center"
        justifyContent="space-between"
        sx={{ my: 2 }}
      >
        {/* <Checkbox name="remember" label="Remember me" />
        <Link variant="subtitle2" underline="hover">
          Forgot password?
        </Link> */}
      </Stack>
      <LoadingButton
        fullWidth
        size="large"
        type="submit"
        variant="contained"
        onClick={handleSubmit}
      >
        Login
      </LoadingButton>
      <ErrorModal
        open={errorModalOpen}
        message={errorMessage}
        onClose={handleCloseErrorModal}
      />
    </>
  );
}
