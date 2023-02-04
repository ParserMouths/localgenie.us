import React from "react";
import "../Styles/tag.scss";

export default function Tag({ title }) {
  return (
    <div className={`tag ${title === "Open" ? "bg-green" : "bg-red"}`}>
      <p>{title}</p>
    </div>
  );
}
