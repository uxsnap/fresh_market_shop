import { MakeOrderItem } from "@/types";
import client from "../client";

type Body = {
  products: MakeOrderItem[];
};

export const makeOrder = (body: Body): Promise<{ data: { uid: string } }> => {
  return client.post("/orders", body);
};

makeOrder.queryKey = "makeOrder";
