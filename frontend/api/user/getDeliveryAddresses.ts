import { User, DeliveryAddress } from "@/types";
import client from "../client";

export const getDeliveryAddresses = () => {
  return client.get<DeliveryAddress[]>("/user/addresses");
};

getDeliveryAddresses.queryKey = "getDeliveryAddresses";
