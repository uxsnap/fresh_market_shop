import client from "../client";

type Body = {
  uid: string;
  photos: string[];
};

export const deleteRecipePhotos = (body: Body) => {
  return client.post("/recipes/photos/delete", body);
};

deleteRecipePhotos.queryKey = "deleteRecipePhotos";
