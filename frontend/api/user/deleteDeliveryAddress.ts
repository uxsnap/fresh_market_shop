import client from "../client";

type Body = {
  addressUid: string;
  apartment?: number;
  floor?: number;
  entrance?: number;
  code?: number;
};

export const deleteDeliveryAddress = (addressUid: string) => {
  return client.post(`/user/addresses/delete`, { addressUid });
};

deleteDeliveryAddress.queryKey = "deleteDeliveryAddress";
