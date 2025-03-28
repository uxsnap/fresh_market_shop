"use client";

import cn from "classnames";
import { Box, Container, Group, Title, Text, Stack } from "@mantine/core";

import styles from "./CartMain.module.css";
import { CartList } from "@/components/CartList";
import { useCartStore } from "@/store";
import { RemoveAll } from "../RemoveAll";
import { formatDuration } from "@/utils";
import { Delivery } from "@/components/icons/Delivery";

type Props = {
  empty?: boolean;
};

export const CartMain = ({ empty = true }: Props) => {
  const storeItems = useCartStore((state) => state.items);
  const delivery = useCartStore((s) => s.delivery);

  const items = Object.values(storeItems).map((item) => item);

  const calculateDelivery = () => {
    if (!delivery || empty) {
      return null;
    }

    const time = formatDuration(delivery.time / 1000);

    return (
      <Text fw="bold" fz={18} c="accent.1">
        Доставка около {!time ? "5 минут" : time}
      </Text>
    );
  };

  const calculatedDelivery = calculateDelivery();

  return (
    <Container
      p={0}
      className={cn(styles.root, empty && styles.empty)}
      m={0}
      mt={20}
    >
      <Group mb={20} align="center" justify="space-between">
        <Stack gap={8}>
          <Title order={1} c="accent.0">
            Корзина
          </Title>

          {!empty && calculatedDelivery && (
            <Group align="center" gap={12} className={styles.delivery}>
              <Delivery size={24} fill="var(--mantine-color-accent-1)" />

              {calculatedDelivery}
            </Group>
          )}
        </Stack>

        <Box visibleFrom="md">{!empty && <RemoveAll />}</Box>
      </Group>

      {!empty && <CartList items={items} />}
    </Container>
  );
};
