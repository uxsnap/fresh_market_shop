import client from "../client";

type Body = {
  name: string;
  ccal: number;
  cookingTime: number;
};

export const createRecipe = (body: Body) => {
  return client.post<{ uid: string }>("/recipes", body);
};

createRecipe.queryKey = "createRecipe";
