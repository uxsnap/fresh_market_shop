import client from "../client";

type Body = {
  uid: string;
  name: string;
  ccal: number;
  cookingTime: number;
};

export const editRecipe = (body: Body) => {
  return client.put("/recipes", body);
};

editRecipe.queryKey = "editRecipe";
