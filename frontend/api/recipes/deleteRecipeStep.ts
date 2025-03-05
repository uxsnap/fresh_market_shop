import client from "../client";

type Body = {
  uid: string;
  step: number;
};

export const deleteRecipeStep = (body: Body) => {
  return client.post("/recipes/steps/delete", body);
};

deleteRecipeStep.queryKey = "deleteRecipeStep";
