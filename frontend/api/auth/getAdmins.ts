import { Admin } from "@/types";
import client from "../client";

export const getAdmins = () => {
  return client.get<Admin[]>("/auth/admin");
};

getAdmins.queryKey = "getAdmins";
