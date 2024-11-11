import { User } from "@/types";
import client from "../client";

export const getUser = () => {
  return client.get<User>("/user");
};

getUser.queryKey = "getUser";
