import client from "../client";

type Body = {
  email: string;
  password: string;
};

export const createAdminUser = (body: Body) => {
  return client.post("/auth/admin", body);
};

createAdminUser.queryKey = "createAdminUser";
