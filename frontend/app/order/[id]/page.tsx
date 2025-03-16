"use client";

import { BackToCatalog } from "@/components/BackToCatalog";
import { PaymentBlock } from "@/components/PaymentBlock";
import { CART_MAIN_HEIGHT } from "@/constants";
import { Box, Group } from "@mantine/core";
import { useEffect } from "react";

import styles from "./order.module.css";
import { PayButton } from "@/components/pages/cart/PayButton";
import { useParams, useRouter } from "next/navigation";
import { useAuthStore } from "@/store/auth";
import { OrderMain } from "@/components/pages/order/OrderMain";
import { useQuery } from "@tanstack/react-query";
import { getOrder } from "@/api/order/getOrder";

export default function OrderPage() {
  const router = useRouter();
  const logged = useAuthStore((s) => s.logged);

  const { id } = useParams();

  useEffect(() => {
    if (logged === undefined) {
      return;
    }

    if (!logged) {
      router.push("/");
    }
  }, [logged]);

  const { data, isFetching } = useQuery({
    queryFn: () => getOrder(id + ""),
    queryKey: [getOrder.queryKey],
  });

  return (
    <>
      <Box className={styles.root}>
        <BackToCatalog />

        <Box mih={CART_MAIN_HEIGHT} pos="relative">
          <Group
            className={styles.group}
            wrap="nowrap"
            align="flex-start"
            justify="space-between"
            w="100%"
          >
            <OrderMain />

            <Box className={styles.paymentBlock} w="100%">
              <PaymentBlock buttonText="Оплатить" />
            </Box>
          </Group>
        </Box>
      </Box>

      <PayButton />
    </>
  );
}
