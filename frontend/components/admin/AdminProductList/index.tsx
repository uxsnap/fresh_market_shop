import { getProducts } from "@/api/products/getProducts";
import { ItemCard } from "@/components/ItemCard";
import { useAdminStore } from "@/store/admin";
import { ProductItem } from "@/types";
import { convertProductToProductItem } from "@/utils";
import { Group, LoadingOverlay, Pagination, Stack } from "@mantine/core";
import { useWindowScroll } from "@mantine/hooks";
import { useQuery } from "@tanstack/react-query";
import { useEffect, useState } from "react";
import { Filters } from "./Filters";

const PRODUCTS_LIMIT = 30;

export const AdminProductList = () => {
  const [activePage, setPage] = useState(1);
  const [_, scrollTo] = useWindowScroll();

  const [filters, setFilters] = useState({ name: "" });

  const setModalOpen = useAdminStore((s) => s.setModalOpen);
  const setProductItem = useAdminStore((s) => s.setProductItem);

  const { data, isLoading } = useQuery({
    queryKey: [getProducts.queryKey, activePage, filters],
    queryFn: () =>
      getProducts({
        page: activePage,
        limit: PRODUCTS_LIMIT,
        ...filters,
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

  const handleEdit = (item: ProductItem) => {
    setProductItem(item);
    setModalOpen(true);
  };

  return (
    <Stack gap={20} pos="relative">
      <LoadingOverlay
        visible={isLoading}
        zIndex={1}
        overlayProps={{ radius: "sm", blur: 2 }}
        loaderProps={{ color: "primary.0", type: "bars" }}
      />

      <Filters setFilters={setFilters} />

      <Group mt={12} gap={12} justify="center">
        {(data?.products ?? []).map((item, ind) => (
          <ItemCard
            onExtended={() => handleEdit(item)}
            editable
            item={item}
            key={ind}
          />
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
