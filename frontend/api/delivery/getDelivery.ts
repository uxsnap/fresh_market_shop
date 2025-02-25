import { DeliveryData } from "@/types";
import client from "../client";

type Body = {
  orderUid: string;
  deliveryAddressUid: string;
};

export const getDelivery = (body: Body) => {
  return client.post<DeliveryData>("/delivery/calculate", body);
};

getDelivery.queryKey = "getDelivery";
