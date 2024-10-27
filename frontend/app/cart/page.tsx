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

// TODO: Figure out right height proportions (CART_MAIN_HEIGHT)
export default function CartPage() {
  const items = useCartStore((s) => s.items);

  const [empty, setEmpty] = useState(false);

  useEffect(() => {
    setEmpty(!Object.keys(items).length);
  }, [items]);

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
                  onClick={() => console.log("ordered")}
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
