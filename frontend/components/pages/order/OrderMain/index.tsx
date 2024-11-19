"use client";

import { CreditCardItemList } from "@/components/CreditCardItemList";
import { Box, Stack, Title } from "@mantine/core";

import styles from "./OrderMain.module.css";

export const OrderMain = () => {
  return (
    <Box className={styles.root} mt={12}>
      <Title order={1} c="accent.0">
        Оформление заказа
      </Title>

      <Stack mt={24} gap={20}>
        <Stack gap={20}>
          <Title order={2} c="accent.0">
            Способ оплаты
          </Title>

          <CreditCardItemList />
        </Stack>
      </Stack>
    </Box>
  );
};
