import { useState, useEffect } from "react";
import NavBar from "./Components/Navbar.jsx";
import Home from "./Pages/Home.jsx";
import Results from "./Pages/Results.jsx";
import VendorLanding from "./Pages/VendorLanding.jsx";
import VendorHome from "./Pages/VendorHome.jsx";
import { useAuth } from "./hooks/useAuth.js";

import { TransitionGroup, CSSTransition } from "react-transition-group";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
  Redirect,
  useLocation,
  useParams,
} from "react-router-dom";

import "./App.scss";
import RequireAuth from "./Components/RequireAuth.jsx";
// import { RequireAuth } from "./Components/RequireAuth.jsx";
//<FeaturedCard img={require('./Assets/fruits.png')} title="Fruits" description="juicy asf" />
//<VendorCard img={require('./Assets/vendor-1.png')} title="Gavin Belson's Sandwich" description="After working in tech, i finally decided to sell the sandwiches."/>
//<VendorLanding />
//<Results />

//<TransitionGroup>
//<CSSTransition
//key={location.pathname}
//classNames="fade"
//timeout={300}
//>
//<Switch location={location}>
//<Route path="/hsl/:h/:s/:l" children={<HSL />} />
//<Route path="/rgb/:r/:g/:b" children={<RGB />} />
//</Switch>
//</CSSTransition>
//</TransitionGroup>

// const { auth, setAuth } = useAuth();
// setAuth({
//   isVendor: localStorage.getItem("isVendor"),
//   user: localStorage.getItem("user"),
// });
function App() {
  let location = useLocation();
  const userId = localStorage.getItem("userId");
  return (
    <div className="App">
      <TransitionGroup>
        <CSSTransition key={location.pathname} classNames="page" timeout={300}>
          <Switch location={location}>
            <Route exact path="/" children={<Start />} />
            <Route
              path="/user/home"
              children={<RequireAuth children={<Home className="home" />} />}
            />
            <Route path="/user/results/:searchQuery" children={<Results />} />
            <Route path="/vendor/home" children={<VendorHome />} />
            {/* </Route> */}
          </Switch>
        </CSSTransition>
      </TransitionGroup>
      {userId && <NavBar className="navbar" />}
    </div>
  );
}
export default App;
