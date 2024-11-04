import axios from "axios";

type Body = {
  email: string;
  password: string;
};

export const loginUser = (body: Body) => {
  return axios.post(`${process.env.NEXT_PUBLIC_API_PROXY_BASE_URL}/login`, body);
};

loginUser.queryKey = "loginUser";
