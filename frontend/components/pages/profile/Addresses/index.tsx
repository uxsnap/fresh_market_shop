import { AddressItemList } from "@/components/AddressItemList";
import { Stack, Title } from "@mantine/core";
import { memo } from "react";

type Props = {
  offsetScrollbars?: boolean;
};

export const Addresses = memo(({ offsetScrollbars = true }: Props) => (
  <Stack gap={16}>
    <Title c="accent.0" order={3}>
      Адрес доставки
    </Title>

    <AddressItemList offsetScrollbars={offsetScrollbars} />
  </Stack>
));
