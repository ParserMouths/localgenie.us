import React from "react";
import { useHistory } from "react-router-dom";

export default function Logout() {
  localStorage.clear();
  let history = useHistory();
  history.push("/");
  return;
}
