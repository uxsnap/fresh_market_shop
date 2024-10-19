import { getRecommendations } from "@/api/recommendations/getRecommendations";
import { ItemList } from "@/components/ItemList";
import { ProductItem } from "@/types";
import { useQuery } from "@tanstack/react-query";

export const YouMayLike = () => {
  const { data, isFetching } = useQuery({
    queryKey: [getRecommendations.queryKey],
    queryFn: () => getRecommendations({ with_photos: true }),
    select(data): ProductItem[] {
      return data.data.map((item) => ({
        imgs: (item.photos ?? []).map((p) => `${process.env.NEXT_PUBLIC_API}/${p.path}`),
        price: item.product.price,
        name: item.product.name,
        info: `${item.product.weight}грамм/${item.product.ccal}ккал`,
      }));
    },
  });

  return <ItemList items={data} isFetching={isFetching} />;
};
