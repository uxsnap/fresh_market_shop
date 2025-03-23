"use client";

import { Cart } from "@/components/icons/Cart";
import { useAuthStore } from "@/store/auth";
import { useCartStore } from "@/store/cart";
import { Button } from "@mantine/core";
import { useRouter } from "next/navigation";

export const CartButton = () => {
  const router = useRouter();

  const logged = useAuthStore((s) => s.logged);
  const setModalOpen = useAuthStore((s) => s.setModalOpen);

  const price = useCartStore((s) => s.getFullPrice());
  const itemsPrice = useCartStore((s) => s.getItemsPrice());

  const handleCartClick = () => {
    if (logged) {
      return router.push("/cart");
    }

    setModalOpen("login");
  };

  return (
    <Button
      onClick={handleCartClick}
      variant={itemsPrice ? "accent" : "accent-reverse"}
      h={38}
      leftSection={
        <Cart fill={itemsPrice ? "white" : "var(--mantine-color-accent-0)"} />
      }
    >
      {itemsPrice ? `${price + 10} руб.` : "Корзина"}
    </Button>
  );
};
