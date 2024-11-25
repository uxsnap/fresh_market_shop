import client from "../client";

type Body = {
  addressUid: string;
  apartment?: number;
  floor?: number;
  entrance?: number;
  code?: number;
};

export const addDeliveryAddress = (body: Body) => {
  return client.post("/user/addresses", body);
};

addDeliveryAddress.queryKey = "addDeliveryAddress";
