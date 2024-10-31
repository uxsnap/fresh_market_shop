import { getRecipeProducts } from "@/api/recipes/getRecipeProducts";
import { ItemList } from "@/components/ItemList";
import { ProductItem } from "@/types";
import { convertProductToProductItem } from "@/utils";
import { useQuery } from "@tanstack/react-query";

type Props = {
  uid: string;
};

export const RecipeProducts = ({ uid }: Props) => {
  const { data, isFetching } = useQuery({
    queryFn: () => getRecipeProducts(uid),
    queryKey: [getRecipeProducts.queryKey, uid],
    select(data): ProductItem[] {
      return data.data.map(convertProductToProductItem);
    },
  });

  return  <ItemList noTitle items={data} isFetching={isFetching} />
};