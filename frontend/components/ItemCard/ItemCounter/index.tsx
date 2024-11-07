import { Counter } from "@/components/Counter";
import { useCartStore } from "@/store";
import { ProductItem } from "@/types";
import { Box, Button } from "@mantine/core";
import { PropsWithChildren } from "react";

type Props = {
  item?: ProductItem;
};

export const ItemCounter = ({
  item,
  children = "Добавить",
}: PropsWithChildren<Props>) => {
  if (!item) {
    return null;
  }

  const { incCartItem, decCartItem, addCartItem, getCount } = useCartStore();

  const count = getCount(item.id);

  return (
    <Box mt={8}>
      {count === 0 ? (
        <Button
          fz={12}
          w="100%"
          onClick={() => addCartItem(item)}
          variant="accent"
        >
          {children}
        </Button>
      ) : (
        <Counter
          count={count}
          onDecrement={() => decCartItem(item.id)}
          onIncrement={() => incCartItem(item.id)}
        />
      )}
    </Box>
  );
};
