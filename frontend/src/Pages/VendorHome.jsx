import React, { useState, useEffect } from "react";
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import { Carousel } from "react-responsive-carousel";
import Mybutton from "../Components/Button.jsx";
import { withRouter, Link, useHistory } from "react-router-dom";

import "../Styles/vendorhome.scss";
import SliderButton from "../Components/SliderButton.jsx";
import axios from "../utils/axios/axios.js";
import Loader from "../Components/Loader.jsx";
import authHeader from "../utils/axios/auth-header.js";
import Tag from "../Components/Tag.jsx";
// import MyButton from "../Components/Button.jsx";

const dummyData = {
  imgs: [
    require("../Assets/vendor-1.png"),
    require("../Assets/vendor-2.png"),
    require("../Assets/vendor-3.png"),
  ],
  title: "Russ Hennman's Samosas",
  description:
    "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
};

const editHandler = (_) => {
  console.log("edit");
};

export default function VendorHome(props) {
  let history = useHistory();
  const [editing, setEditing] = useState(false);
  const [loading, setLoading] = useState(false);
  const editHandler = (_) => setEditing(!editing);

  const paymentHandler = (_) => {
    if (editing) {
      const newUrl = window.prompt("Enter New Payment URL");
    } else console.log("regular");
  };
  const [slider, setSlider] = useState(true);

  // useEffect(() => {

  // }, [slider]);
  const sliderHandler = () => {
    setSlider(!slider);
    (async () => {
      setLoading(true);
      const res = await axios.post(
        `/stall/update/${localStorage.getItem("stallId")}`,
        {
          is_open: slider ? 1 : 0,
          latitude: localStorage.getItem("latitude"),
          longitude: localStorage.getItem("longitude"),
        },
        {
          headers: authHeader(),
        }
      );
      setLoading(false);
      if (res.status == 200) alert("Operation is successful!");
      else alert("Error updating Info");
    })();
  };

  const logoutHandler = () => {
    console.log("logging out");
    localStorage.clear();
    history.push("/");
  };

  return (
    <div className={props.className}>
      {loading && <Loader />}
      <div className="vendor-home-page">
        <Mybutton className="edit-btn" onClick={editHandler}>
          {editing ? (
            "Done"
          ) : (
            <>
              <i className="far fa-pen"></i> &nbsp; Edit
            </>
          )}
        </Mybutton>
        <div className="carousel-wrapper">
          <Carousel
            showArrows={true}
            showThumbs={false}
            dynamicHeight={true}
            infiniteLoop={true}
            autoPlay={true}
          >
            {dummyData["imgs"].map((img, i) => (
              <img src={img} key={i} />
            ))}
          </Carousel>
        </div>

        <div className="vendor-content">
          <hr />
          <h2 className={editing ? "edit" : ""} contentEditable={editing}>
            {dummyData["title"]}
          </h2>
          <div className="btns">
            <Mybutton outlined={true} className="btn-x">
              {" "}
              Navigate{" "}
            </Mybutton>
            <Mybutton className="btn-x" onClick={paymentHandler}>
              {editing ? "Edit Payment Link" : "Pay Merchant"}
            </Mybutton>
          </div>
          <div className="vendor-slider">
            <h3>Close</h3>
            <SliderButton sliderHandler={sliderHandler} />
            <h3>Open</h3>
          </div>
          {/* <Tag title={}/> */}
          <h3>Offering</h3>
          <p className={editing ? "edit" : ""} contentEditable={editing}>
            {dummyData["description"]}
          </p>

          <h3>About Vendor</h3>
          <p className={editing ? "edit" : ""} contentEditable={editing}>
            {dummyData["description"]}
          </p>
        </div>
        <div className="vendor-btn">
          <Mybutton onClick={logoutHandler}>Logout</Mybutton>
        </div>
        {/* <Tags name="Not Active" /> */}
      </div>
    </div>
  );
}
