import { AddressItemList } from "@/components/AddressItemList";
import { Stack, Title } from "@mantine/core";
import { memo } from "react";

export const Addresses = memo(() => (
  <Stack gap={16}>
    <Title c="accent.0" order={2}>
      Адрес доставки
    </Title>

    <AddressItemList />
  </Stack>
));
