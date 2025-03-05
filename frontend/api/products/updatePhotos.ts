import client from "../client";

export const updatePhotos = (form: FormData) => {
  return client.post("/recipes/photos", form);
};

updatePhotos.queryKey = "updatePhotos";
