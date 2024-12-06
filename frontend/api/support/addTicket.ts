import { Topic } from "@/types";
import client from "../client";

type Body = {
  topicUid: string;
  title: string;
  fromEmail: string;
  fromPhone?: string;
  description: string;
};

export const addTicket = (body: Body) => {
  return client.post("/support/tickets", body);
};

addTicket.queryKey = "addTicket";
