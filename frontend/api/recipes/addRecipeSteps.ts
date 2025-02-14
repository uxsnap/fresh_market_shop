import { RecipeStep } from "@/types";
import client from "../client";

type Body = {
  uid: string;
  rSteps: RecipeStep[];
};

export const addRecipeSteps = (body: Body) => {
  return client.post("/recipes/steps", body);
};

addRecipeSteps.queryKey = "addRecipeSteps";
