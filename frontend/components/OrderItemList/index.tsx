import { Stack, Title } from "@mantine/core";
import { OrderItem, Props as OrderItemProps } from "../OrderItem";

type Props = {
  orders?: OrderItemProps[];
};

export const OrderItemList = ({
  orders = Array.from({ length: 5 }),
}: Props) => {
  return (
    <Stack gap={12}>
      {orders.map((order, ind) => (
        <OrderItem key={ind} {...order} />
      ))}
    </Stack>
  );
};
