import React, { useEffect, useState } from "react";
import MyButton from "../Components/Button.jsx";
import "../Styles/login.scss";
import axios from "../utils/axios/axios";
import { LOGIN_URL } from "../utils/constants.js";
import { useAuth } from "../hooks/useAuth.js";
import { Link, Redirect, useHistory } from "react-router-dom";
import { userGetLocation } from "../utils/getLocation.js";
import Loader from "../Components/Loader.jsx";
export default function Login() {
  let history = useHistory();
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);
  // const { auth, setAuth } = useAuth();

  const isVendor = localStorage.getItem("isVendor");

  const handleSubmit = async (e) => {
    e.preventDefault();
    console.log(username, password);
    setLoading(true);
    try {
      const res = await axios.post(
        LOGIN_URL,
        JSON.stringify({
          Username: username,
          Email: email,
          Password: password,
        })
      );
      console.log(res.data);
      localStorage.setItem("userId", res.data.user_id);
      localStorage.setItem("token", res.data.token);
      localStorage.setItem("stallId", res.data.stall_id);
      await userGetLocation();
      // <Redirect to="user/login" />;
    } catch (err) {
      console.log(err);
    }
    setLoading(false);
    if (isVendor) history.push("/createStore");
    else history.push("/user/home");
  };
  return (
    <div className="login">
      {loading ? (
        <Loader />
      ) : (
        <section className="form-wrapper">
          <h1>Sign In</h1>
          <form onSubmit={handleSubmit} className="form">
            <div className="form-input">
              <input
                type="text"
                id="username"
                placeholder="Username"
                autoComplete="on"
                onChange={(e) => setUsername(e.target.value)}
                value={username}
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
                type="password"
                id="password"
                placeholder="Password"
                onChange={(e) => setPassword(e.target.value)}
                value={password}
                required
              />
            </div>

            <MyButton style={{ width: "100vw", textAlign: "center" }}>
              &nbsp; Sign In
            </MyButton>
          </form>
          <p>
            Need an Account?
            <span>
              {/*put router link here*/}
              <Link to="/signup" style={{ textDecoration: "none" }}>
                &nbsp; Sign Up
              </Link>
            </span>
          </p>
        </section>
      )}
    </div>
  );
}
