import client from "../client";

type Body = {
  addressUid: string;
  apartment?: number;
  floor?: number;
  entrance?: number;
  code?: number;
};

export const addUserAddress = (body: Body) => {
  return client.post("/user/addresses", body);
};

addUserAddress.queryKey = "addUserAddress";
