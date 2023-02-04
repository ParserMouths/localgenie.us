import React, { useState } from "react";
import MyButton from "../Components/Button.jsx";
import "../Styles/login.scss";
import axios from "../utils/axios/axios";
import { SIGNUP_URL } from "../utils/constants.js";
import { Link, Redirect, useHistory } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";
import { userGetLocation } from "../utils/getLocation.js";

export default function SignUp(props) {
  // const { auth } = useAuth();

  const [username, setUsername] = useState("");
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [number, setNumber] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  let history = useHistory();
  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const [latitude, longitude] = await userGetLocation();
      const payload = {
        is_vendor: localStorage.getItem("isVendor"),
        username: username,
        first_name: firstName,
        last_name: lastName,
        latitude: latitude,
        longitude: longitude,
        phone_no: number,
        email: email,
        password: password,
        subscription: localStorage.getItem("subscription"),
      };
      const res = await axios.post(SIGNUP_URL, payload);
      console.log(res);
      history.push("/login");
      // <Redirect to="/login" />;
      // console.log("Hi");
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <div className="">
      <section className="form-wrapper">
        <h1>SignUp</h1>
        <form onSubmit={handleSubmit} className="form">
          <div className="form-input">
            <input
              type="text"
              id="username"
              placeholder="Username"
              autoComplete="off"
              onChange={(e) => setUsername(e.target.value)}
              value={username}
              required
            />

            <input
              type="text"
              id="firstname"
              placeholder="Firstname"
              autoComplete="off"
              onChange={(e) => setFirstName(e.target.value)}
              value={firstName}
              required
            />

            <input
              type="text"
              id="lastname"
              placeholder="Lastname"
              autoComplete="off"
              onChange={(e) => setLastName(e.target.value)}
              value={lastName}
              required
            />

            <input
              type="email"
              id="email"
              placeholder="Email"
              autoComplete="on"
              onChange={(e) => setEmail(e.target.value)}
              value={email}
              required
            />

            <input
              type="number"
              id="number"
              placeholder="Phone Number"
              autoComplete="on"
              onChange={(e) => setNumber(e.target.value)}
              value={number}
              required
            />

            <input
              type="password"
              id="password"
              placeholder="Password"
              onChange={(e) => setPassword(e.target.value)}
              value={password}
              required
            />

            <input
              type="password"
              id="confirmPassword"
              placeholder="Confirm Password"
              onChange={(e) => setConfirmPassword(e.target.value)}
              value={confirmPassword}
              required
            />
          </div>
          <MyButton style={{ width: "100vw", textAlign: "center" }}>
            Sign Up
          </MyButton>
        </form>
        <p>
          Already registered?
          <span>
            {/*put router link here*/}
            <Link to="/login" style={{ "text-decoration": "none" }}>
              Sign In
            </Link>
          </span>
        </p>
      </section>
    </div>
  );
}
