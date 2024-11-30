"use client";

import { BackToCatalog } from "@/components/BackToCatalog";
import { PaymentBlock } from "@/components/PaymentBlock";
import { CART_MAIN_HEIGHT } from "@/constants";
import { useCartStore } from "@/store";
import { Box, Group } from "@mantine/core";
import { useEffect, useState } from "react";

import styles from "./order.module.css";
import { PayButton } from "@/components/pages/cart/PayButton";
import { useParams, useRouter } from "next/navigation";
import { useAuthStore } from "@/store/auth";
import { OrderMain } from "@/components/pages/order/OrderMain";
import { useQuery } from "@tanstack/react-query";
import { getOrder } from "@/api/order/getOrder";

export default function OrderPage() {
  const router = useRouter();
  const items = useCartStore((s) => s.items);
  const logged = useAuthStore((s) => s.logged);

  const { id } = useParams();

  const [empty, setEmpty] = useState(false);

  useEffect(() => {
    setEmpty(!Object.keys(items).length);
  }, [items]);

  useEffect(() => {
    if (logged === undefined) {
      return;
    }

    if (!logged || empty) {
      router.push("/");
    }
  }, [logged, empty]);

  const { data, isFetching } = useQuery({
    queryFn: () => getOrder(id + ""),
    queryKey: [getOrder.queryKey],
  });

  console.log(data);

  return (
    <>
      <Box className={styles.root}>
        <BackToCatalog empty={empty} />

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
