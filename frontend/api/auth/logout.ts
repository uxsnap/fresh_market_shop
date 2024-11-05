import client from "../client";

export const logoutUser = () => {
  return client.post("/logout");
};

logoutUser.queryKey = "logoutUser";
