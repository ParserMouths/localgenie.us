import React, { useState, useRef } from "react";
import MyButton from "../Components/Button.jsx";
import "../Styles/createstore.scss";
import axios from "../utils/axios/axios";
import { Link, Redirect, useHistory } from "react-router-dom";
import { CREATE_STALL_URL } from "../utils/constants.js";
import authHeader from "../utils/axios/auth-header.js";
import Loader from "../Components/Loader.jsx";
import { BASE_URL } from "../utils/constants.js";

export default function CreateStore(props) {
  // const { auth } = useAuth();
  const isVendor = localStorage.getItem("isVendor");
  const userId = localStorage.getItem("userId");

  const [stallName, setstallName] = useState("");
  const [aboutVendor, setAboutVendor] = useState("");
  const [offerings, setOfferings] = useState("");
  const fileRef = useRef();
  const [loading, setLoading] = useState(false);

  let history = useHistory();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
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
      const res = await fetch(`${BASE_URL}/stall/new`, {
        method: "POST",
        body: formData,
        headers: {
          ...authHeader(),
        },
      }).then((r) => r.json());
      console.log(res);
      localStorage.setItem("stallId", res.stall_id);
      history.push("/user/home");
    } catch (err) {
      console.log(err);
    }
    setLoading(false);
  };

  return (
    <div className="">
      {loading ? (
        <Loader />
      ) : (
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
                required
              />

              <textarea
                id="aboutVendor"
                placeholder="AboutVendor"
                autoComplete="off"
                onChange={(e) => setAboutVendor(e.target.value)}
                required
                rows="4"
              />

              <textarea
                id="offerings"
                placeholder="Offerings"
                autoComplete="off"
                onChange={(e) => setOfferings(e.target.value)}
                rows="4"
                required
              />

              <input
                type="file"
                id="Files"
                name="Files"
                multiple
                ref={fileRef}
              />
            </div>
            <MyButton style={{ width: "100vw", textAlign: "center" }}>
              Create Stall
            </MyButton>
          </form>
        </section>
      )}
    </div>
  );
}
