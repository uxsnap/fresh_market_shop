import client from "../client";

export const deleteRecipe = (uid: string) => {
  return client.post("/recipes/delete", { uid });
};

deleteRecipe.queryKey = "deleteRecipe";
