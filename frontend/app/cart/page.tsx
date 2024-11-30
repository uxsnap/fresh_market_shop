"use client";

import { BackToCatalog } from "@/components/BackToCatalog";
import { HugeIconText } from "@/components/HugeIconText";
import { CartMain } from "@/components/pages/cart/CartMain";
import { PaymentBlock } from "@/components/PaymentBlock";
import { YouMayLike } from "@/components/YouMayLike";
import { CART_MAIN_HEIGHT } from "@/constants";
import { useCartStore } from "@/store";
import { Box, Group } from "@mantine/core";
import { useEffect, useState } from "react";

import styles from "./cart.module.css";
import { PayButton } from "@/components/pages/cart/PayButton";
import { useMutation } from "@tanstack/react-query";
import { useRouter } from "next/navigation";
import { makeOrder } from "@/api/order/makeOrder";
import { showErrorNotification } from "@/utils";
import { AxiosError } from "axios";
import { useAuthStore } from "@/store/auth";

// TODO: Figure out right height proportions (CART_MAIN_HEIGHT)
export default function CartPage() {
  const router = useRouter();
  const items = useCartStore((s) => s.items);
  const logged = useAuthStore((s) => s.logged);
  const removeAllItems = useCartStore((s) => s.removeAllItems);

  const [empty, setEmpty] = useState(false);

  useEffect(() => {
    setEmpty(!Object.keys(items).length);
  }, [items]);

  useEffect(() => {
    if (logged === undefined) {
      return;
    }

    if (!logged) {
      router.push("/");
    }
  }, [logged]);

  const mutation = useMutation({
    mutationFn: makeOrder,
    onSuccess: (data) => {
      router.push(`/order/${data.uid}`);
      removeAllItems();
    },
    onError: (error: AxiosError<any>) => {
      showErrorNotification(error);
    },
  });

  const handleCreateOrder = () => {
    mutation.mutate({
      products: Object.entries(items).map(([id, item]) => ({
        productUid: id,
        count: item.count,
      })),
    });
  };

  const renderEmpty = () => {
    if (!empty) return;

    return (
      <HugeIconText center type="sad">
        В корзине нет товаров
      </HugeIconText>
    );
  };

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
            <CartMain empty={empty} />
            {!empty && (
              <Box className={styles.paymentBlock} w="100%">
                <PaymentBlock
                  buttonText="Оформить заказ"
                  onClick={handleCreateOrder}
                />
              </Box>
            )}
          </Group>

          {renderEmpty()}
        </Box>

        <Box mt={20}>
          <YouMayLike />
        </Box>
      </Box>

      <PayButton />
    </>
  );
}
