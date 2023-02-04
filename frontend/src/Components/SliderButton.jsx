import React from "react";
import "../Styles/sliderbutton.scss";

export default function SliderButton({ sliderHandler }) {
  return (
    <div class="toggle">
      <label class="switch">
        <input type="checkbox" onChange={sliderHandler} />
        <span class="slider round"></span>
      </label>
    </div>
  );
}
