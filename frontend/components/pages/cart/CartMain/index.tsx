"use client";

import cn from "classnames";
import { Box, Container, Group, Title } from "@mantine/core";

import styles from "./CartMain.module.css";
import { CartList } from "@/components/CartList";
import { useCartStore } from "@/store";
import { RemoveAll } from "../RemoveAll";

type Props = {
  empty?: boolean;
};

export const CartMain = ({ empty = true }: Props) => {
  const storeItems = useCartStore((state) => state.items);

  const items = Object.values(storeItems).map((item) => item);

  return (
    <Container
      p={0}
      className={cn(styles.root, empty && styles.empty)}
      m={0}
      mt={20}
    >
      <Group mb={20} align="center" justify="space-between">
        <Title order={1} c="accent.0">
          Корзина
        </Title>

        <Box visibleFrom="md">{!empty && <RemoveAll />}</Box>
      </Group>

      {!empty && <CartList items={items} />}
    </Container>
  );
};
