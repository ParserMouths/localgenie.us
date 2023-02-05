import { useState, useEffect } from "react";
import NavBar from "./Components/Navbar.jsx";
import Home from "./Pages/Home.jsx";
import Results from "./Pages/Results.jsx";
import VendorLanding from "./Pages/VendorLanding.jsx";
import VendorHome from "./Pages/VendorHome.jsx";
import SignUp from "./Pages/SignUp.jsx";
import CreateStore from "./Pages/CreateStore.jsx";
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
import Login from "./Pages/Login.jsx";
import Start from "./Pages/Start.jsx";
import RequireAuth from "./Components/RequireAuth.jsx";
import Profile from "./Pages/Profile.jsx";
import Logout from "./Pages/Logout.jsx";
import { BASE_URL } from "./utils/constants.js";

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
  console.log(location);

  useEffect((_) => {
    function urlBase64ToUint8Array(base64String) {
      var padding = "=".repeat((4 - (base64String.length % 4)) % 4);
      var base64 = (base64String + padding)
        .replace(/\-/g, "+")
        .replace(/_/g, "/");

      var rawData = window.atob(base64);
      var outputArray = new Uint8Array(rawData.length);

      for (var i = 0; i < rawData.length; ++i) {
        outputArray[i] = rawData.charCodeAt(i);
      }
      return outputArray;
    }
    if ("serviceWorker" in navigator) {
      window.addEventListener("load", function () {
        navigator.serviceWorker
          .register("serviceWorker.js")
          .then(function (registration) {
            // Use the PushManager to get the user's subscription to the push service.
            return registration.pushManager
              .getSubscription()
              .then(async function (subscription) {
                // If a subscription was found, return it.
                if (subscription) {
                  return subscription;
                }

                // Get the server's public key
                const response = await fetch(`${BASE_URL}/notification/key`);
                const result = await response.json();
                console.log(result);
                // Chrome doesn't accept the base64-encoded (string) vapidPublicKey yet
                // urlBase64ToUint8Array() is defined in /tools.js
                const convertedVapidKey = urlBase64ToUint8Array(
                  result["vapid_key"]
                );

                // Otherwise, subscribe the user (userVisibleOnly allows to specify that we don't plan to
                // send notifications that don't have a visible effect for the user).
                return registration.pushManager.subscribe({
                  userVisibleOnly: true,
                  applicationServerKey: convertedVapidKey,
                });
              });
          })
          .then(function (subscription) {
            localStorage.setItem("subscription", JSON.stringify(subscription));
            // Send the subscription details to the server using the Fetch API.
            //   fetch('http://localhost:6969/notification/register', {
            // 	method: 'post',
            // 	headers: {
            // 	  'Content-type': 'application/json'
            // 	},
            // 	body: JSON.stringify({
            // 	  subscription: subscription
            // 	}),
            //   });
          });
      });
    }
  }, []);

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
            <Route path="/signup" children={<SignUp />} />
            <Route path="/login" children={<Login />} />
            <Route
              path="/createStore"
              children={<RequireAuth children={<CreateStore />} />}
            />
            <Route
              path="/profile"
              children={<RequireAuth children={<Profile />} />}
            />
            <Route
              path="/logout"
              children={<RequireAuth children={<Logout />} />}
            />
          </Switch>
        </CSSTransition>
      </TransitionGroup>
      {(location?.pathname.includes("/user") ||
        location?.pathname.includes("/vendor")) &&
        userId && <NavBar className="navbar" />}
    </div>
  );
}
//<Route path="/user/home" children={<Home className='home'/>} />
//<Route path="/user/results" children={<Results />} />
//<Route path="/vendor/home" children={<VendorLanding />} />
export default App;
