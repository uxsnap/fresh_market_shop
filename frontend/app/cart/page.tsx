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
    <Box pt={22} px={8} mx="auto" maw={1520}>
      <BackToCatalog />

      <Box mih={CART_MAIN_HEIGHT} pos="relative">
        <Group
          wrap="nowrap"
          gap={60}
          align="flex-start"
          justify="space-between"
        >
          <CartMain empty={empty} />

          <Box w="100%" maw={362}>
            <PaymentBlock
              text="Оформить заказ"
              onClick={() => console.log("ordered")}
            />
          </Box>
        </Group>

        {renderEmpty()}
      </Box>

      <Box mt={20}>
        <YouMayLike />
      </Box>
    </Box>
  );
}
