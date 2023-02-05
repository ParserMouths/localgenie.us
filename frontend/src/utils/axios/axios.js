import axios from "axios";
import { BASE_URL } from "../constants";

export default axios.create({
  //Base URL of Backend.
  baseURL: BASE_URL,
});
