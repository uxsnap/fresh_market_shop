import { getOrderProducts } from "@/api/order/getOrderProducts";
import { ItemList } from "@/components/ItemList";
import { ProductItem } from "@/types";
import { convertProductToProductItem } from "@/utils";
import { useQuery } from "@tanstack/react-query";
import { memo } from "react";

export const YouAlreadyOrdered = memo(() => {
  const { data, isFetching, isError } = useQuery({
    queryKey: [getOrderProducts.queryKey],
    queryFn: getOrderProducts,
    select(data): ProductItem[] {
      return data.data.map(convertProductToProductItem);
    },
    retry: 0,
  });

  if (!isFetching && (!data?.length || isError)) {
    return null;
  }

  return (
    <ItemList title="Вы уже заказывали" items={data} isFetching={isFetching} />
  );
});
