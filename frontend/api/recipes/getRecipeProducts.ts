import { ProductWithPhotos } from "@/types";
import client from "../client";

export const getRecipeProducts = (uid: string) => {
  return client.get<ProductWithPhotos[]>(`/recipes/${uid}/products`);
};

getRecipeProducts.queryKey = "getRecipeProducts";
