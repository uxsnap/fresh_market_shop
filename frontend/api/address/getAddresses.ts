import { Address } from "@/types";
import client from "../client";

export const getAddresses = (cityUid: string, name: string) => {
  return client.get<Address[]>(`/addresses/${cityUid}`, {
    params: { name },
  });
};

getAddresses.queryKey = "getAddresses";
