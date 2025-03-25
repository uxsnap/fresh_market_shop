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
import { useMutation, useQuery } from "@tanstack/react-query";
import { getOrder } from "@/api/order/getOrder";
import { makePayment } from "@/api/order/makePayment";
import { useCartStore } from "@/store";
import { useOrderStore } from "@/store/order";
import { showErrorNotification } from "@/utils";
import { AxiosError } from "axios";

export default function OrderPage() {
  const router = useRouter();
  const logged = useAuthStore((s) => s.logged);
  const delivery = useCartStore((s) => s.delivery);
  const creditCard = useOrderStore((s) => s.creditCard);

  const { id } = useParams();

  useEffect(() => {
    if (logged === undefined) {
      return;
    }

    if (!logged) {
      router.push("/");
    }
  }, [logged]);

  useEffect(() => {});

  const { data } = useQuery({
    queryFn: () => getOrder(id + ""),
    queryKey: [getOrder.queryKey],
  });

  const { mutate } = useMutation({
    mutationFn: makePayment,
    mutationKey: [makePayment.queryKey],
    onSuccess: () => {
      router.push("/payment_complete");
    },
    onError: (error: AxiosError<any>) => {
      showErrorNotification(error);
    },
  });

  const isDisabled = !delivery || !creditCard || !data?.data || !id;

  const handlePayment = () => {
    if (isDisabled) {
      return;
    }

    mutate({
      orderUid: id + "",
      cardUid: creditCard?.uid,
      sum: data?.data.sum,
      currency: "RUB",
    });
  };

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
              <PaymentBlock
                onClick={handlePayment}
                disabled={isDisabled}
                buttonText="Оплатить"
                price={data?.data.sum}
              />
            </Box>
          </Group>
        </Box>
      </Box>

      <PayButton>Оплатить</PayButton>
    </>
  );
}
