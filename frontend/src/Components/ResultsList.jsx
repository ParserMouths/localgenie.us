import React from "react";
import VendorCard from "./VendorCard.jsx";

import "../Styles/resultslist.scss";

export default function ResultsList(props) {
  return (
    <div className={props.className}>
      <div className="results-list-wrapper">
        {props.data.map((d, i) => (
          <VendorCard
            width={"80vw"}
            key={i}
            className="vendor-card"
            data={d}
            animate={true}
            to={{ search: `stall_id=${d["stall_id"]}` }}
            setCurrentMarker={props.setCurrentMarker}
          />
        ))}
      </div>
    </div>
  );
}
//<VendorCard className="vendor-card" img={require('../Assets/vendor-1.png')} title="Gavin Belson's Sandwich" description="After working in tech, i finally decided to sell the sandwiches."/>
//<VendorCard className="vendor-card" img={require('../Assets/vendor-3.png')} title="Russ Hennman's Samosas" description="After working in tech, i finally decided to sell the Samosas."/>
//<VendorCard className="vendor-card" img={require('../Assets/vendor-2.png')} title="Monica's Fruit Bowl" description="After working in tech, i finally decided to sell the Fruits."/>
