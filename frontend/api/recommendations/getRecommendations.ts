import { ProductsWithTotal } from "@/types";
import client from "../client";

type Params = {
  user_uid?: string;
  category_uid?: string;
  with_photos?: boolean;
  limit?: number;
};

export const getRecommendations = (params?: Params) => {
  return client.get<ProductsWithTotal>("/recommendations", {
    params,
  });
};

getRecommendations.queryKey = "getRecommendations";
