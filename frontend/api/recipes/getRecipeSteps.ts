import { RecipeStep } from "@/types";
import client from "../client";

export const getRecipeSteps = (uid: string) => {
  return client.get<RecipeStep[]>(`/recipes/${uid}/steps`);
};

getRecipeSteps.queryKey = "getRecipeSteps";
