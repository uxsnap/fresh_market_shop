import { AddressItemList } from "@/components/AddressItemList";
import { ShadowBox } from "@/components/ShadowBox";
import { Stack, Title } from "@mantine/core";

export const Addresses = () => {
  return (
    <ShadowBox w="100%" mah={312} style={{ overflowY: "auto" }}>
      <Stack gap={16} p={12}>
        <Title c="accent.0" order={3}>
          Адреса доставки
        </Title>

        <AddressItemList />
      </Stack>
    </ShadowBox>
  );
};
