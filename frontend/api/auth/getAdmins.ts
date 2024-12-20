import { Admin } from "@/types";
import client from "../client";

export const getAdmins = () => {
  return client.get<{ admins: Admin[] }>("/auth/admins");
};

getAdmins.queryKey = "getAdmins";
