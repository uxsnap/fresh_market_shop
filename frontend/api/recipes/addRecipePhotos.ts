import client from "../client";

export const addRecipePhotos = (form: FormData) => {
  return client.post("/recipes/photos", form);
};

addRecipePhotos.queryKey = "addRecipePhotos";
