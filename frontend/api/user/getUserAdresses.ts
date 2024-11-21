import { User, UserAddress } from "@/types";
import client from "../client";

export const getUserAddresses = () => {
  return client.get<UserAddress[]>("/user/addresses");
};

getUserAddresses.queryKey = "getUserAddresses";
