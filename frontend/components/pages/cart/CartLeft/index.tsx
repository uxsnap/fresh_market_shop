"use client";

import { Button, Container, Group, Title } from "@mantine/core";

import styles from "./CartLeft.module.css";
import { Trash } from "@/components/icons/Trash";
import { useCartStore } from "@/store";

export const CartLeft = () => {
  const state = useCartStore();

  return (
    <Container className={styles.root} m={0} mt={20}>
      <Group align="center" justify="space-between">
        <Title order={1} c="accent.0">
          Корзина
        </Title>

        <Button
          mih={38}
          variant="accent-reverse"
          leftSection={<Trash fill="var(--mantine-color-accent-0)" />}
        >
          Очистить корзину
        </Button>
      </Group>
    </Container>
    
  );
};
