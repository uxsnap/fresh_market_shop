import { JwtData } from "@/types";
import client from "../client";

export const getUserInfo = () => {
  return client.get<JwtData>("/userInfo");
};

getUserInfo.queryKey = "getUserInfo";
