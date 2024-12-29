import client from "../client";

type Body = {
  uid: string;
  price: number;
  name: string;
  weight: number;
  ccal: number;
  description: string;
  categoryUid: string;
};

export const editProduct = (body: Body) => {
  return client.put("/products", body);
};

editProduct.queryKey = "editProduct";
