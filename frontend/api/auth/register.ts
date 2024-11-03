import client from "../client";

type Body = {
  email: string;
  password: string;
};

export const registerUser = (body: Body) => {
  return client.post("/auth/register", body);
};

registerUser.queryKey = "registerUser";
