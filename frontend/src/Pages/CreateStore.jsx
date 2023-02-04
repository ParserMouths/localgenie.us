import React, { useState, useRef } from "react";
import MyButton from "../Components/Button.jsx";
import "../Styles/createstore.scss";
import axios from "../utils/axios/axios";
import { Link, Redirect, useHistory } from "react-router-dom";
import { CREATE_STALL_URL } from "../utils/constants.js";
import authHeader from "../utils/axios/auth-header.js";

export default function CreateStore(props) {
  // const { auth } = useAuth();
  const isVendor = localStorage.getItem("isVendor");
  const userId = localStorage.getItem("userId");

  const [stallName, setstallName] = useState("");
  const [aboutVendor, setAboutVendor] = useState("");
  const [offerings, setOfferings] = useState("");
  const fileRef = useRef();

  let history = useHistory();

  const handleSubmit = async (e) => {
    e.preventDefault();
    const files = fileRef.current;
    const formData = new FormData();
    for (const file of files.files) {
      formData.append("files", file);
    }
    formData.append("owner", userId);
    formData.append("stall_name", stallName);
    formData.append("latitude", localStorage.getItem("latitude"));
    formData.append("longitude", localStorage.getItem("longitude"));
    formData.append("offering", offerings);
    formData.append("about_vendor", aboutVendor);

    try {
      console.log(authHeader());
      const res = await fetch("http://127.0.0.1:6969/stall/new", {
        method: "POST",
        body: formData,
        headers: {
          ...authHeader(),
        },
      });
      history.push("/user/home");
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <div className="">
      <section className="form-wrapper">
        <h1>Setup Shop</h1>
        <form
          onSubmit={handleSubmit}
          className="form"
          encType="multipart/form-data"
        >
          <div className="form-input">
            <input
              type="text"
              id="stallName"
              placeholder="StallName"
              autoComplete="off"
              onChange={(e) => setstallName(e.target.value)}
              value={stallName}
              required
            />

            <textarea
              id="aboutVendor"
              placeholder="AboutVendor"
              autoComplete="off"
              onChange={(e) => setAboutVendor(e.target.value)}
              required
              rows="4"
            >
              {aboutVendor}
            </textarea>

            <textarea
              id="offerings"
              placeholder="Offerings"
              autoComplete="off"
              onChange={(e) => setOfferings(e.target.value)}
              rows="4"
              required
            >
              {offerings}
            </textarea>

            <input type="file" id="Files" name="Files" multiple ref={fileRef} />
          </div>
          <MyButton style={{ width: "100vw", textAlign: "center" }}>
            Create Stall
          </MyButton>
        </form>
      </section>
    </div>
  );
}
