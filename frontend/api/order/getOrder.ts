import { Order } from "@/types";
import client from "../client";

export const getOrder = (orderId: string) => {
  return client.get<Order[]>(`/orders`, { params: { orderId } });
};

getOrder.queryKey = "getOrder";
