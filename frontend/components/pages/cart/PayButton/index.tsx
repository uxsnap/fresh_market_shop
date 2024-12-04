"use client";

import { Delivery } from "@/components/icons/Delivery";
import { useCartStore } from "@/store";
import { Box, Button, Group, Text } from "@mantine/core";

import styles from "./PayButton.module.css";
import { useEffect, useState } from "react";
import { useMutation } from "@tanstack/react-query";
import { makeOrder } from "@/api/order/makeOrder";
import { useRouter } from "next/navigation";
import { showErrorNotification } from "@/utils";
import { AxiosError } from "axios";

export const PayButton = () => {
  const [curPrice, setCurPrice] = useState<number>(0);
  const router = useRouter();

  const price = useCartStore((s) => s.getFullPrice());
  const items = useCartStore((s) => s.items);
  const removeAllItems = useCartStore((s) => s.removeAllItems);

  const mutation = useMutation({
    mutationFn: makeOrder,
    onSuccess: (data) => {
      router.push(`/order/${data.data.uid}`);
    },
    onError: (error: AxiosError<any>) => {
      showErrorNotification(error);
    },
  });

  useEffect(() => {
    setCurPrice(price);
  }, [price]);

  const handleCreateOrder = () => {
    mutation.mutate({
      products: Object.entries(items).map(([id, item]) => ({
        productUid: id,
        count: item.count,
      })),
    });
  };

  if (!curPrice) {
    return null;
  }

  return (
    <Box hiddenFrom="md" className={styles.root}>
      <Button onClick={handleCreateOrder} w="100%" variant="accent" h={40}>
        <Group gap={16} align="center">
          <Text fw="bold" fz={18}>
            Оформить заказ
          </Text>

          <Group gap={8} align="center">
            <Text fw="bold" fz={18}>
              {curPrice} ₽
            </Text>
            <Delivery size={24} fill="var(--mantine-color-bg-2)" />
          </Group>
        </Group>
      </Button>
    </Box>
  );
};
