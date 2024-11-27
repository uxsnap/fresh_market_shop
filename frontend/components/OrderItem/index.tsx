import { Group, Stack, Title, Text, Image } from "@mantine/core";
import { OrderStatus, OrderWithProducts } from "@/types";
import { dayJs } from "@/utils";

export type Props = OrderWithProducts;

const mapStatusToText: Record<OrderStatus, string> = {
  new: "Создан",
  paid: "Оплачен",
  in_progress: "В работе |",
  done: "Доставлен",
};

export const OrderItem = ({ order, products }: Props) => (
  <Stack w="100%" p={12} bg="bg.1">
    <Group w="100%" justify="space-between" align="flex-start">
      <Group gap={12}>
        <Stack gap={0}>
          <Title order={4} c="accent.0">
            Заказ #{order.num}
          </Title>

          <Text fz={12} c="accent.3">
            {mapStatusToText[order.status]}{" "}
            {dayJs(order.updatedAt).format("DD.MM.YYYY в HH:mm")}
          </Text>
        </Stack>
      </Group>

      <Title order={4} c="accent.0">
        {order.sum} Руб
      </Title>
    </Group>

    <Group gap={12}>
      {products.map((product) => (
        <Image
          radius={8}
          key={product.productUid}
          mah={60}
          src={product.photos[0].path}
        />
      ))}
    </Group>
  </Stack>
);
