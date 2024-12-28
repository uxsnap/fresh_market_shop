import client from "../client";

export const deleteProduct = (uid: string) => {
  return client.post("/products/delete", { uid });
};

deleteProduct.queryKey = "deleteProduct";
