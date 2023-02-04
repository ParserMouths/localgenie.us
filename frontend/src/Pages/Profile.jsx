import React, { useEffect, useState } from "react";
import "../Styles/profile.scss";
import MyButton from "../Components/Button.jsx";
import { Link } from "react-router-dom";

export default function Profile() {
  const username = "Raghavendra";
  const email = "raghav@gmail.com";
  return (
    <div className="profile-wrapper">
      <div className="profile-card">
        <img src={require("../Assets/profile.png")} />
        <div className="title-description">
          <p> Username: {username}</p>
          <p> Email: {email}</p>
        </div>
      </div>
      <div className="profile-btn">
        <Link to="/logout">
          <MyButton>&nbsp; Logout</MyButton>
        </Link>
      </div>
    </div>
  );
}
