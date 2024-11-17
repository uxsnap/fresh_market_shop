import { Address } from "@/types";
import client from "../client";

export const getAddresses = (
  cityUid: string,
  name: string,
  house_number: string
) => {
  return client.get<Address[]>(`/addresses/${cityUid}`, {
    params: { name, house_number },
  });
};

getAddresses.queryKey = "getAddresses";
