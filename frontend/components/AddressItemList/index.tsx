"use client";

import { Button, ScrollArea, Stack } from "@mantine/core";
import { AddressItem } from "../AddressItem";
import { Plus } from "../icons/Plus";
import { getAddress } from "@/utils";
import { UserAddress } from "@/types";

type Props = {
  activeAddress?: UserAddress;
  setActiveAddress: (v: UserAddress) => void;
  onAdd: () => void;
  items?: UserAddress[];
};

export const AddressItemList = ({
  items,
  activeAddress,
  setActiveAddress,
  onAdd,
}: Props) => (
  <Stack gap={12}>
    <Button
      mr={12}
      onClick={onAdd}
      mih={48}
      variant="dashed"
      leftSection={<Plus fill="var(--mantine-color-accent-0)" />}
    >
      Добавить
    </Button>

    <ScrollArea h={250} offsetScrollbars>
      <Stack gap={12}>
        {items?.map((address) => (
          <AddressItem
            onSelect={() => setActiveAddress(address)}
            active={address.addressUid === activeAddress?.addressUid}
            key={address.addressUid}
          >
            {getAddress(address)}
          </AddressItem>
        ))}
      </Stack>
    </ScrollArea>
  </Stack>
);
