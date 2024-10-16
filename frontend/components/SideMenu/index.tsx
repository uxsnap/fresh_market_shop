import { Container, Flex, Stack, Title, Text } from "@mantine/core";
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
  "Хлебобулочные изделия": Bread,
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
  const { data } = useQuery({ queryKey: [], queryFn: getCategories });

  console.log(data);

  return (
    <Container m={0} p={0}>
      <Stack gap={20}>
        <Title c="accent.0" order={2}>
          Каталог
        </Title>

        <Stack gap={12}>
          {/* {menu.map(({ label, icon: Icon }) => (
            <Flex py={4} px={12} className={styles.item} key={label} gap={10}>
              <Icon />

              <Text c="accent.0">{label}</Text>
            </Flex> */}
          {/* ))} */}
        </Stack>
      </Stack>
    </Container>
  );
};
