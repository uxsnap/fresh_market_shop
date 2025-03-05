"use client";

import { Box, Stack, Title } from "@mantine/core";

import { useQuery } from "@tanstack/react-query";
import React from "react";
import { getCategories } from "@/api/categories/getCategories";
import { CategoryItem } from "../CategoryItem";
import { SkeletLoader } from "../SkeletLoader";
import { useParams, useRouter } from "next/navigation";

export const SideMenu = () => {
  const router = useRouter();
  const params = useParams<{ category_uid?: string }>();

  const { data, isFetching } = useQuery({
    queryKey: [getCategories.queryKey],
    queryFn: getCategories,
    refetchOnWindowFocus: false,
    staleTime: Infinity,
  });

  const active =
    data?.data.find((c) => c.uid === params.category_uid)?.name ?? "Главная";

  const renderLoader = () => <SkeletLoader l={8} />;

  const renderData = () => (
    <>
      <CategoryItem
        active={active === "Главная"}
        onClick={() => router.push(`/`)}
      >
        Главная
      </CategoryItem>
      {data?.data.map(({ uid, name }) => (
        <CategoryItem
          active={active === name}
          onClick={() => router.push(`/products/${uid}`)}
          key={uid}
        >
          {name}
        </CategoryItem>
      ))}
    </>
  );

  return (
    <Box>
      <Stack gap={20}>
        <Title c="accent.0" order={2}>
          Каталог
        </Title>

        <Stack gap={12}>
          {isFetching && renderLoader()}

          {!isFetching && renderData()}
        </Stack>
      </Stack>
    </Box>
  );
};
