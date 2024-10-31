import { Category, Recipe } from "@/types";
import client from "../client";

export const getRecipes = () => {
  return client.get<Recipe[]>("/recipes");
};

getRecipes.queryKey = "getRecipes";
