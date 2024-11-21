import client from "../client";

type Body = {
  latitude: number;
  longitude: number;
  cityName: string;
  streetName: string;
  houseNumber: string;
  building: number;
  floor: number;
  apartment: number;
  code: number;
};

export const addUserAddress = (body: Body) => {
  return client.post("/user/addresses", body);
};

addUserAddress.queryKey = "addUserAddress";
