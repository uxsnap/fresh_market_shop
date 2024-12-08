import { ISupportItem } from "@/types";
import client from "../client";

export const getTickets = () => {
  return client.get<ISupportItem[]>("/support/tickets");
};

getTickets.queryKey = "getTickets";
