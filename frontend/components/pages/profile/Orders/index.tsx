import { OrderItemList } from "@/components/OrderItemList";
import { ShadowBox } from "@/components/ShadowBox";
import { Stack, Title } from "@mantine/core";

export const Orders = () => {
  return (
    <ShadowBox p={12} mah={480} style={{ zIndex: 1, overflowY: "auto" }}>
      <Stack gap={16}>
        <Title c="accent.0" order={3}>
          История заказов
        </Title>

        <OrderItemList />
      </Stack>
    </ShadowBox>
  );
};
