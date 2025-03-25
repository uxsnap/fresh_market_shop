"use client";

import { CreditCardItemList } from "@/components/CreditCardItemList";
import { Box, Stack, Title, Text, Group } from "@mantine/core";

import styles from "./OrderMain.module.css";
import { Addresses } from "../../profile/Addresses";
import { useCartStore } from "@/store";
import { Delivery } from "@/components/icons/Delivery";
import { formatDuration } from "@/utils";

export const OrderMain = () => {
  const delivery = useCartStore((s) => s.delivery);

  const calculateDelivery = () => {
    if (!delivery) {
      return null;
    }

    const time = formatDuration(delivery.time / 1000);

    return (
      <Text fw="bold" fz={18} c="accent.1">
        Доставка около {!time ? "5 минут" : time}
      </Text>
    );
  };

  return (
    <Box className={styles.root} mt={12}>
      <Title order={1} c="accent.0">
        Оформление заказа
      </Title>

      <Group mt={8} align="center" gap={12}>
        <Delivery size={24} fill="var(--mantine-color-accent-1)" />

        {calculateDelivery()}
      </Group>

      <Stack className={styles.main} gap={20}>
        <Stack gap={20}>
          <Title order={2} c="accent.0">
            Способ оплаты
          </Title>

          <CreditCardItemList />

          <Addresses />
        </Stack>
      </Stack>
    </Box>
  );
};
