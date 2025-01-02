import { ProductsWithTotal } from "@/types";
import client from "../client";

type Params = {
  limit?: number;
  page?: number;
  name?: string;
};

export const getProducts = ({ limit = 30, page = 1, name = "" }: Params) => {
  return client.get<ProductsWithTotal>("/products", {
    params: {
      limit,
      page,
      with_counts: true,
      with_photos: true,
      name,
    },
  });
};

getProducts.queryKey = "getProducts";
