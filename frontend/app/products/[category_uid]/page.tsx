"use client";

import { getCategories } from "@/api/categories/getCategories";
import { getProductsByCategory } from "@/api/categories/getProductsByCategory";
import { ItemList } from "@/components/ItemList";
import { ProductItem } from "@/types";
import { convertProductToProductItem } from "@/utils";
import { Container } from "@mantine/core";
import { useQuery } from "@tanstack/react-query";
import { useParams } from "next/navigation";

export default function ProductPage() {
  const params = useParams<{ category_uid?: string }>();

  const { data: categories } = useQuery({
    queryKey: [getCategories.queryKey],
    queryFn: getCategories,
    refetchOnWindowFocus: false,
    staleTime: Infinity,
  });

  const { data, isFetching } = useQuery({
    queryKey: [getProductsByCategory.queryKey, params.category_uid],
    queryFn: () =>
      getProductsByCategory({
        category_uid: params.category_uid ?? "",
        with_photos: true,
      }),
    select(data): ProductItem[] {
      return data.data.map(convertProductToProductItem);
    },
  });

  const title =
    categories?.data.find((c) => c.uid === params.category_uid)?.name ??
    "Категория";

  return (
    <Container p={8} m={0} miw="100%">
      <ItemList
        title={title}
        scroll={false}
        isFetching={isFetching}
        items={data}
      />
    </Container>
  );
}
