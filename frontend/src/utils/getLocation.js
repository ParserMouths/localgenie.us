import { useEffect, useState } from "react";
import axios from "axios";
export const userGetLocation = () => {
  return new Promise((res, rej) => {
    navigator.geolocation.getCurrentPosition(async ({ coords }) => {
      try {
        console.log(coords);
        localStorage.setItem("latitude", coords.latitude);
        localStorage.setItem("longitude", coords.longitude);
        const { data } = await axios.get(
          `https://api.opencagedata.com/geocode/v1/json?key=c7215cdf8efb48cd8b5fd90aeb5328e1&q=${coords.latitude}+${coords.longitude}&pretty=1&no_annotations=1`
        );
        localStorage.setItem(
          "city_town",
          data.results[0].components.city ||
            data.results[0].components.town ||
            data.results[0].components.state_district
        );
        res([coords.latitude, coords.longitude]);
      } catch (err) {
        console.log(err);
        rej(err);
      }
    });
  });
};
