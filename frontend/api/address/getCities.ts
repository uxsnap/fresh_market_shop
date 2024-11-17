import { City } from "@/types";
import client from "../client";

export const getCities = () => {
  return client.get<City[]>(`/addresses/cities`);
};

getCities.queryKey = "getCities";
