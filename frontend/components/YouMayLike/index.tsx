import { getRecommendations } from "@/api/recommendations/getRecommendations";
import { ItemList } from "@/components/ItemList";
import { ProductItem } from "@/types";
import { convertProductToProductItem } from "@/utils";
import { useQuery } from "@tanstack/react-query";
import { memo } from "react";

export const YouMayLike = memo(() => {
  const { data, isFetching } = useQuery({
    queryKey: [getRecommendations.queryKey],
    queryFn: () => getRecommendations({ with_photos: true, limit: 50 }),
    select(data): ProductItem[] {
      return data.data.map(convertProductToProductItem);
    },
  });

  return (
    <ItemList
      title="Вам может понравиться"
      items={data}
      isFetching={isFetching}
    />
  );
});
