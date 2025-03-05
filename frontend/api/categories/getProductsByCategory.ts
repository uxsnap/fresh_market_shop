import { ProductWithPhotos } from "@/types";
import client from "../client";

type Params = {
  category_uid: string;
  with_photos?: boolean;
};

export const getProductsByCategory = ({ category_uid, ...rest }: Params) => {
  return client.get<{ products: ProductWithPhotos[] }>(
    `/categories/${category_uid}/products`,
    {
      params: rest,
    }
  );
};

getProductsByCategory.queryKey = "getProductsByCategory";
