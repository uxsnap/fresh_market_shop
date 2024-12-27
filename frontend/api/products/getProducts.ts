import { ProductsWithTotal } from "@/types";
import client from "../client";

type Params = {
  limit?: number;
  page?: number;
};

export const getProducts = ({ limit = 30, page = 1 }: Params) => {
  return client.get<ProductsWithTotal>("/products", {
    params: {
      limit,
      page,
      with_counts: true,
      with_photos: true,
    },
  });
};

getProducts.queryKey = "getProducts";
