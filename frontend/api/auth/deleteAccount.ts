import client from "../client";

export const deleteAccount = () => {
  return client.delete("/user/delete");
};

deleteAccount.queryKey = "deleteAccount";
