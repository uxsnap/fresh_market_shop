import client from "../client";

type Body = {
  uid: string;
  photoUid: string;
};

export const deleteProductPhoto = (body: Body) => {
  return client.post("/products/photos/delete", body);
};

deleteProductPhoto.queryKey = "deleteProductPhoto";
