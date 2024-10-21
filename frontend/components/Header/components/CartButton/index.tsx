"use client";

import { Cart } from "@/components/icons/Cart";
import { useCartStore } from "@/store/cart";
import { Button } from "@mantine/core";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

export const CartButton = () => {
  const router = useRouter();

  const [curPrice, setCurPrice] = useState<number>(0);

  const price = useCartStore((s) => s.getFullPrice());

  useEffect(() => {
    setCurPrice(price);
  }, [price]);

  return (
    <Button
      onClick={() => router.push("/cart")}
      visibleFrom="sm"
      variant={curPrice ? "accent" : "accent-reverse"}
      h={38}
      leftSection={
        <Cart fill={curPrice ? "white" : "var(--mantine-color-accent-0)"} />
      }
    >
      {curPrice ? `${curPrice} руб.` : "Корзина"}
    </Button>
  );
};
