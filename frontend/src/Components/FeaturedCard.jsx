import React from "react";
import { Link } from "react-router-dom";
import "../Styles/featuredcard.scss";

export default function FeaturedCard(props) {
  return (
    <div className={props.className}>
      <Link
        to={`/user/results/${props.title}`}
        style={{ textDecoration: "None" }}
      >
        <div className="wrapper">
          <img src={props.img} />
          <h3> {props.title} </h3>
          <p> {props.description} </p>
        </div>
      </Link>
    </div>
  );
}
