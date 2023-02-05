import React from "react";
import "../Styles/button.scss";

export default function MyButton(props) {
  console.log(props.onClick);
  return (
    <button
      onClick={props.onClick}
      className={`${props.outlined ? "btn-outlined" : "btn"} ${
        props.className
      }`}
    >
      {props.children}
    </button>
  );
}
