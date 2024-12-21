import client from "../client";

type Body = {
  uid: string;
};

export const deleteAccount = (body?: Body) => {
  return client.post("/user/delete", body ?? {});
};

deleteAccount.queryKey = "deleteAccount";
