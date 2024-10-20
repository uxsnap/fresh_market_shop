import { getRecommendations } from "@/api/recommendations/getRecommendations";
import { ItemList } from "@/components/ItemList";
import { ProductItem } from "@/types";
import { convertProductToProductItem } from "@/utils";
import { useQuery } from "@tanstack/react-query";

export const YouMayLike = () => {
  const { data, isFetching } = useQuery({
    queryKey: [getRecommendations.queryKey],
    queryFn: () => getRecommendations({ with_photos: true }),
    select(data): ProductItem[] {
      return data.data.map(convertProductToProductItem);
    },
  });

  return <ItemList items={data} isFetching={isFetching} />;
};
