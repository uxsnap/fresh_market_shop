"use client";

import { BackToCatalog } from "@/components/BackToCatalog";
import { HugeIconText } from "@/components/HugeIconText";
import { MainBox } from "@/components/MainBox";
import { CartLeft } from "@/components/pages/cart/CartLeft";
import { useCartStore } from "@/store";
import { useEffect, useState } from "react";

export default function CartPage() {
  const items = useCartStore((s) => s.items);

  const [empty, setEmpty] = useState(false);

  useEffect(() => {
    setEmpty(!Object.keys(items).length);
  }, [items]);

  const renderMain = () => <></>;

  const renderEmpty = () => {
    return (
      <HugeIconText center type="sad">
        В корзине нет товаров
      </HugeIconText>
    );
  };

  return (
    <MainBox pt={22} pos="relative">
      <BackToCatalog />

      <CartLeft empty={empty} />

      {!empty ? renderMain() : renderEmpty()}
    </MainBox>
  );
}
