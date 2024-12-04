import { OrderWithProducts } from "@/types";
import client from "../client";

export const getOrdersHistory = () => {
  return client.get<OrderWithProducts[]>(`/orders/history`);
};

getOrdersHistory.queryKey = "getOrdersHistory";
