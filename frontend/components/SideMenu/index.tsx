import { Container, Stack, Title } from "@mantine/core";

import { useQuery } from "@tanstack/react-query";
import React from "react";
import { getCategories } from "@/api/categories/getCategories";
import { CategoryItem } from "../CategoryItem";
import { SkeletLoader } from "../SkeletLoader";

export const SideMenu = () => {
  const { data, isFetching } = useQuery({
    queryKey: [getCategories.queryKey],
    queryFn: getCategories,
    refetchOnWindowFocus: false,
  });

  const renderLoader = () => <SkeletLoader l={8} />;

  const renderData = () =>
    data?.data.map(({ name }) => (
      <CategoryItem key={name}>{name}</CategoryItem>
    ));

  return (
    <Container m={0} p={0}>
      <Stack gap={20}>
        <Title c="accent.0" order={2}>
          Каталог
        </Title>

        <Stack gap={12}>
          {isFetching && renderLoader()}

          {!isFetching && renderData()}
        </Stack>
      </Stack>
    </Container>
  );
};
