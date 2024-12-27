import { getProducts } from "@/api/products/getProducts";
import { ItemCard } from "@/components/ItemCard";
import { ProductItem } from "@/types";
import { convertProductToProductItem } from "@/utils";
import { Group, LoadingOverlay, Pagination, Stack } from "@mantine/core";
import { useWindowScroll } from "@mantine/hooks";
import { useQuery } from "@tanstack/react-query";
import { useEffect, useState } from "react";

const PRODUCTS_LIMIT = 30;

export const AdminProductList = () => {
  const [activePage, setPage] = useState(1);
  const [_, scrollTo] = useWindowScroll();

  const { data, isFetching } = useQuery({
    queryKey: [getProducts.queryKey, activePage],
    queryFn: () =>
      getProducts({
        page: activePage,
        limit: PRODUCTS_LIMIT,
      }),
    select(data): { total: number; products: ProductItem[] } {
      return {
        total: data.data.total,
        products: data.data.products.map(convertProductToProductItem),
      };
    },
  });

  useEffect(() => {
    scrollTo({ y: 0 });
  }, [activePage]);

  return (
    <Stack gap={20} pos="relative" justify="center" align="center">
      <LoadingOverlay
        visible={isFetching}
        zIndex={1}
        overlayProps={{ radius: "sm", blur: 2 }}
        loaderProps={{ color: "primary.0", type: "bars" }}
      />

      <Group gap={12} justify="center">
        {(data?.products ?? []).map((item, ind) => (
          <ItemCard item={item} key={ind} />
        ))}
      </Group>

      {data?.total && (
        <Pagination
          color="accent.0"
          value={activePage}
          onChange={setPage}
          total={Math.ceil(data.total / PRODUCTS_LIMIT)}
        />
      )}
    </Stack>
  );
};
