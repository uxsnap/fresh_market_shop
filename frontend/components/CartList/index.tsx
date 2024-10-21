import { Flex, ScrollArea, Stack } from "@mantine/core";
import { CartItem } from "../CartItem";
import { useCartStore } from "@/store";
import { CartItem as ICartItem } from "@/types";

type Props = {
  items: ICartItem[];
};

export const CartList = ({ items }: Props) => {
  const inc = useCartStore((state) => state.incCartItem);
  const dec = useCartStore((state) => state.decCartItem);
  const remove = useCartStore((state) => state.removeCartItem);

  return (
    <Flex gap={20} mah={432} pos="relative" direction="column">
      <ScrollArea h="100%">
        <Stack gap={8}>
          {items.map((item) => {
            return (
              <CartItem
                {...item.product}
                onDecrement={() => dec(item.product.id)}
                onIncrement={() => inc(item.product.id)}
                onDelete={() => remove(item.product.id)}
                count={item.count}
                key={item.product.id}
              />
            );
          })}
        </Stack>
      </ScrollArea>
    </Flex>
  );
};
