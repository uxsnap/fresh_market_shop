import client from "../client";

type Body = {
  email: string;
  password: string;
};

export const loginUser = (body: Body) => {
  return client.post("/login", body);
};

loginUser.queryKey = "loginUser";
