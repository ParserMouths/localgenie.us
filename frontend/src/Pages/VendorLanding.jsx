import React, { useEffect, useState } from "react";
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import { Carousel } from "react-responsive-carousel";
import "../Styles/vendorlanding.scss";
import Mybutton from "../Components/Button.jsx";
import { withRouter, useLocation } from "react-router-dom";
import { getStore } from "../utils/storyblok.js";
import Loader from "../Components/Loader";
import Tag from "../Components/Tag";

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

function useQuery() {
  const { search } = useLocation();

  return React.useMemo(() => new URLSearchParams(search), [search]);
}

function VendorLanding(props) {
  const [data, setData] = useState(undefined);
  const [loading, setLoading] = useState(true);
  let query = useQuery();

  useEffect((_) => {
    (async () => {
      const data = await getStore(query.get("stall_id"));
      console.log("data", data);
      if (data) setData(data);
      setLoading(false);
      // data["IsOpen"] = 0;
      console.log(data);
    })();
  }, []);

  return (
    <div className={props.className}>
      {loading && <Loader />}
      <div className="vendor-landing-page">
        <Mybutton className="back-btn" onClick={props.history.goBack}>
          <i className="far fa-arrow-left"></i> &nbsp; Back{" "}
        </Mybutton>
        <div className="carousel-wrapper">
          <Carousel
            showArrows={true}
            showThumbs={false}
            dynamicHeight={true}
            infiniteLoop={true}
            autoPlay={true}
          >
            {data?.assets?.map((img, i) => (
              <img src={img} key={i} />
            ))}
          </Carousel>
        </div>
        <div className="vendor-content">
          <hr />
          {!loading && (
            <>
              <h2>{data["stall_name"]}</h2>
              <div className="btns">
                <Mybutton outlined={true} className="btn-x">
                  {" "}
                  Navigate{" "}
                </Mybutton>
                <Mybutton className="btn-x"> Pay Merchant </Mybutton>
              </div>

              <Tag
                title={
                  data?.["is_open"] === 1
                    ? "Open"
                    : "Last Active : " + data?.["LastActive"]?.split(" ")[0]
                }
              />

              <h3>Offering</h3>
              <p>{data["offering"]}</p>

              <h3>About Vendor</h3>
              <p>{data["about_vendor"]}</p>
            </>
          )}
        </div>
      </div>
    </div>
  );
}

export default withRouter(VendorLanding);
