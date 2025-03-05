import { RecipesWithTotal } from "@/types";
import client from "../client";

type Args = {
  page?: number;
  limit?: number;
  name?: string;
};

export const getRecipes = ({ page = 1, limit = 20, name = "" }: Args) => {
  return client.get<RecipesWithTotal>("/recipes", {
    params: { limit, page, name },
  });
};

getRecipes.queryKey = "getRecipes";
