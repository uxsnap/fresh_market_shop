import client from "../client";

export const updatePhotos = (form: FormData) => {
  return client.post("/products/photos", form);
};

updatePhotos.queryKey = "updatePhotos";
