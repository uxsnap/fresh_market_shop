import { ProductsWithCategories } from "@/types";
import client from "../client";

export const search = (name: string) => {
  return client.get<ProductsWithCategories>("/search", {
    params: {
      name,
      limit_on_products: 10,
      limit_on_categories: 10,
      page: 0,
      products_with_count: false,
      products_with_photos: true,
    },
  });
};

search.queryKey = "search";
