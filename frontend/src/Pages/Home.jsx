import React from "react";
import Header from "../Components/Header.jsx";
import VendorList from "../Components/VendorList.jsx";
import FeaturedSection from "../Components/FeaturedSection.jsx";
import MyButton from "../Components/Button.jsx";
import SearchBar from "../Components/SearchBar.jsx";
import VendorLanding from "../Pages/VendorLanding.jsx";
import { storyBlokClient } from "../utils/storyblok.js";
import { TransitionGroup, CSSTransition } from "react-transition-group";

import {
  useParams,
  useRouteMatch,
  Route,
  Switch,
  useLocation,
} from "react-router-dom";

import "../Styles/home.scss";
//<h1 style={{textAlign: 'center', width: '100vw', margin: '2rem 0', lineHeight: '120%'}}> Something awesome comming soon </h1>
//<MyButton style={{width: '100vw', textAlign:'center'}}> &nbsp; Stay Tuned ! &nbsp;</MyButton>
//
function useQuery() {
  const { search } = useLocation();

  return React.useMemo(() => new URLSearchParams(search), [search]);
}

export default function Home(props) {
  let { path, url } = useRouteMatch();
  let location = useLocation();
  let query = useQuery();

  return (
    <div className={props.className}>
      <Header className="header" />
      <SearchBar className="search-bar" />
      <FeaturedSection className="featured" />
      <VendorList className="vendor-list" title="Nearby Stalls" />
      <TransitionGroup>
        <CSSTransition key={query} classNames="slide" timeout={300}>
          {/* sneaky workaround to prevent parent to re render */}
          {query.get("StallID") ? (
            <VendorLanding className="vendor-landing-home" />
          ) : (
            <i />
          )}
        </CSSTransition>
      </TransitionGroup>
    </div>
  );
}
