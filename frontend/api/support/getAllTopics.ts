import { Topic } from "@/types";
import client from "../client";

export const getAllTopics = () => {
  return client.get<Topic[]>("/support/tickets/topics");
};

getAllTopics.queryKey = "getAllTopics";
