import { Recommendations } from "@/types";
import client from "../client";

type Params = {
  user_uid?: string;
  category_uid?: string;
  with_photos?: boolean;
};

export const getRecommendations = (params?: Params) => {
  return client.get<Recommendations[]>("/recommendations", {
    params,
  });
};

getRecommendations.queryKey = "getRecommendations";
