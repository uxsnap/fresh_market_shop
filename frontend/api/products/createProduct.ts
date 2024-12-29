import { ProductItem } from "@/types";
import client from "../client";

type Body = {
  price: number;
  name: string;
  weight: number;
  ccal: number;
  description: string;
  categoryUid: string;
};

export const createProduct = (body: Body) => {
  return client.post("/products", body);
};

createProduct.queryKey = "createProduct";
