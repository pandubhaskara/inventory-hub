import { Route, Navigate } from "react-router-dom";
import PropTypes from 'prop-types';

const PrivateRoute = ({ 
  isAuthenticated
}) => (
  <Route
    render={() =>
      isAuthenticated ? <Navigate to="/dashboard/app" /> : <Navigate to="/login" />
    }
  />
);

PrivateRoute.propTypes = {
  // element: PropTypes.elementType.isRequired,
  isAuthenticated: PropTypes.bool.isRequired,
};

export default PrivateRoute;
