import React from "react";
import { Redirect, Route } from "react-router-dom";

const RequireAuth = ({ children }) => {
  const userId = localStorage.getItem("userId");
  return userId ? children : <Redirect to="/login" />;
};

export default RequireAuth;
