import { Cart } from "@/components/icons/Cart";
import { Button } from "@mantine/core";

export const CartButton = () => {
  return (
    <Button variant="accent-reverse" h={38} leftSection={<Cart />}>
      Корзина
    </Button>
  );
};
