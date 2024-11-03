"use client";

import { Cart } from "@/components/icons/Cart";
import { useAuthStore } from "@/store/auth";
import { useCartStore } from "@/store/cart";
import { Button } from "@mantine/core";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

export const CartButton = () => {
  const router = useRouter();

  const logged = useAuthStore((s) => s.logged);
  const setModalOpen = useAuthStore((s) => s.setModalOpen);

  const [curPrice, setCurPrice] = useState<number>(0);

  const price = useCartStore((s) => s.getFullPrice());

  useEffect(() => {
    setCurPrice(price);
  }, [price]);

  const handleCartClick = () => {
    if (logged) {
      return router.push("/cart");
    }

    setModalOpen("login");
  };

  return (
    <Button
      onClick={handleCartClick}
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
