import { Box, ScrollArea, Stack } from "@mantine/core";
import { CartItem } from "../CartItem";
import { CartItem as ICartItem } from "@/types";
import { CART_MAIN_HEIGHT } from "@/constants";

type Props = {
  items: ICartItem[];
};

export const CartList = ({ items }: Props) => (
  <ScrollArea
    w="100%"
    type="auto"
    h={CART_MAIN_HEIGHT}
    offsetScrollbars
    scrollbars="y"
  >
    <Stack gap={8}>
      {items.map((item) => {
        return <CartItem item={item} key={item.product.id} />;
      })}
    </Stack>
  </ScrollArea>
);
