import React from "react";
import "../Styles/sliderbutton.scss";

export default function SliderButton({ sliderHandler }) {
  return (
    <div className="toggle">
      <label className="switch">
        <input type="checkbox" onChange={sliderHandler} />
        <span className="slider round"></span>
      </label>
    </div>
  );
}
