import { ProductWithPhotos } from "@/types";
import client from "../client";

export const getOrderProducts = () => {
  return client.get<ProductWithPhotos[]>(`/orders/products`);
};

getOrderProducts.queryKey = "getOrderProducts";
