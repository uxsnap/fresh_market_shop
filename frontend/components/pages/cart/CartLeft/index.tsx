"use client";

import cn from "classnames";
import { Button, Container, Group, Title } from "@mantine/core";

import styles from "./CartLeft.module.css";
import { Trash } from "@/components/icons/Trash";
import { CartList } from "@/components/CartList";
import { useCartStore } from "@/store";

type Props = {
  empty?: boolean;
};

export const CartLeft = ({ empty = true }: Props) => {
  const storeItems = useCartStore((state) => state.items);
  const removeAllItems = useCartStore((state) => state.removeAllItems);

  const items = Object.values(storeItems).map((item) => item);

  return (
    <Container className={cn(styles.root, empty && styles.empty)} m={0} mt={20}>
      <Group mb={20} align="center" justify="space-between">
        <Title order={1} c="accent.0">
          Корзина
        </Title>

        {!empty && (
          <Button
            mih={38}
            variant="accent-reverse"
            leftSection={<Trash fill="var(--mantine-color-accent-0)" />}
            onClick={removeAllItems}
          >
            Очистить корзину
          </Button>
        )}
      </Group>

      {!empty && <CartList items={items} />}
    </Container>
  );
};
