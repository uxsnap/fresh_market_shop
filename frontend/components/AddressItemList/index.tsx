"use client";

import { Button, ScrollArea, Stack } from "@mantine/core";
import { AddressItem } from "../AddressItem";
import { Plus } from "../icons/Plus";
import { getAddress } from "@/utils";
import { useMapStore } from "@/store/map";
import { useCallback, useEffect } from "react";
import { getDeliveryAddresses } from "@/api/user/getDeliveryAddresses";
import { useQuery } from "@tanstack/react-query";

type Props = {
  classNames?: {
    button?: string;
  };
};

export const AddressItemList = ({ classNames }: Props) => {
  const deliveryAddress = useMapStore((s) => s.deliveryAddress);
  const setDeliveryAddress = useMapStore((s) => s.setDeliveryAddress);
  const setIsMapOpen = useMapStore((s) => s.setIsMapOpen);

  const { data, isFetched } = useQuery({
    queryFn: getDeliveryAddresses,
    queryKey: [getDeliveryAddresses.queryKey],
  });

  const handleOpenMap = useCallback(() => {
    setIsMapOpen(true);
  }, []);

  useEffect(() => {
    if (data && data.data.length) {
      setDeliveryAddress(data.data[0]);
    }
  }, [isFetched]);

  return (
    <Stack gap={12}>
      <Button
        className={classNames?.button}
        onClick={handleOpenMap}
        mih={48}
        variant="dashed"
        leftSection={<Plus fill="var(--mantine-color-accent-0)" />}
      >
        Добавить
      </Button>

      <ScrollArea h={250} offsetScrollbars>
        <Stack gap={12}>
          {data?.data?.map((address) => (
            <AddressItem
              onSelect={() => setDeliveryAddress(address)}
              active={address.addressUid === deliveryAddress?.addressUid}
              key={address.addressUid}
            >
              {getAddress(address)}
            </AddressItem>
          ))}
        </Stack>
      </ScrollArea>
    </Stack>
  );
};
