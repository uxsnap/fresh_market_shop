import { Recipe } from "@/types";
import client from "../client";

type Body = {
  name: string;
  ccal: number;
  cookingTime: number;
};

export const createRecipe = (body: Body) => {
  return client.post<Recipe>("/recipes", body);
};

createRecipe.queryKey = "createRecipe";
