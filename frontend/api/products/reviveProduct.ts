import client from "../client";

export const reviveProduct = (uid: string) => {
  return client.post("/products/revive", { uid });
};

reviveProduct.queryKey = "reviveProduct";
