import { Container, Stack, Title, Text, Group, Skeleton } from "@mantine/core";
import { Salad } from "../icons/Salad";
import { Bread } from "../icons/Bread";
import { Fish } from "../icons/Fish";
import { Peach } from "../icons/Peach";
import { Tomato } from "../icons/Tomato";
import { Meat } from "../icons/Meat";
import { Oil } from "../icons/Oil";
import { Milk } from "../icons/Milk";
import { Grains } from "../icons/Grains";
import { Bean } from "../icons/Bean";

import styles from "./SideMenu.module.css";
import { useQuery } from "@tanstack/react-query";
import React from "react";
import { getCategories } from "@/api/categories/getCategories";

const mapNameToIcon: Record<string, React.FC> = {
  "Готовая еда": Salad,
  Хлеб: Bread,
  Рыба: Fish,
  Фрукты: Peach,
  Овощи: Tomato,
  "Мясной отдел": Meat,
  Бакалея: Grains,
  Заправка: Oil,
  "Молочные продукты": Milk,
  Топпинги: Bean,
};

export const SideMenu = () => {
  const { data, isFetching } = useQuery({
    queryKey: [getCategories.queryKey],
    queryFn: getCategories,
  });

  const renderLoader = () =>
    Array.from({ length: 8 }, (_, ind) => (
      <Skeleton key={ind} height={32} radius="md" />
    ));

  const renderData = () =>
    data?.data.map(({ name }) => {
      const Icon = mapNameToIcon[name];

      return (
        <Group
          align="center"
          py={4}
          px={12}
          className={styles.item}
          key={name}
          gap={10}
        >
          {Icon ? <Icon /> : ""}

          <Text lh={1} fw={500} fz={18} c="accent.0">
            {name}
          </Text>
        </Group>
      );
    });

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
