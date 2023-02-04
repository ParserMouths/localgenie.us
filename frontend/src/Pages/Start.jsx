import React from "react";
import { Link } from "react-router-dom";
import MyButton from "../Components/Button";
import "../Styles/start.scss";
import { useAuth } from "../hooks/useAuth";
import Logo from "../Components/logo.jsx";

const Start = () => {
  const { setAuth } = useAuth();

  const vendorHandler = () => {
    localStorage.setItem("isVendor", true);
  };
  const clientHandler = () => {
    localStorage.setItem("isVendor", false);
  };
  return (
    <div className="card-start">
      <div className="card-head">
        <Logo style={{ margin: "0.5rem" }} />
        <h3> Are You a Vendor or Customer? </h3>
      </div>
      <div className="card-button">
        <Link to="/login" style={{ width: "calc(50% - 0.5rem)" }}>
          <MyButton
            onClick={vendorHandler}
            outlined={true}
            className="btn-vendor-client"
          >
            Vendor
          </MyButton>
        </Link>
        <Link to="/login" style={{ width: "calc(50% - 0.5rem)" }}>
          <MyButton onClick={clientHandler} className="btn-vendor-client">
            Client
          </MyButton>
        </Link>
      </div>
    </div>
  );
};

export default Start;
