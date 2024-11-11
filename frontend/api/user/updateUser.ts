import client from "../client";

type Body = {
  firstName: string;
  lastName: string;
  birthday?: string;
  email: string;
};

export const updateUser = (body: Body) => {
  return client.put("/users", body);
};

updateUser.queryKey = "updateUser";
