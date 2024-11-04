import axios from "axios";

export default axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_PROXY_BASE_URL,
  timeout: 1000,
});