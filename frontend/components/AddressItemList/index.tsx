"use client";

import { Button, ScrollArea, Stack } from "@mantine/core";
import { AddressItem } from "../AddressItem";
import { Plus } from "../icons/Plus";
import {
  getAddress,
  showErrorNotification,
  showSuccessNotification,
} from "@/utils";
import { useMapStore } from "@/store/map";
import { MouseEventHandler, useEffect } from "react";
import { getDeliveryAddresses } from "@/api/user/getDeliveryAddresses";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { deleteDeliveryAddress } from "@/api/user/deleteDeliveryAddress";

type Props = {
  offsetScrollbars?: boolean;
  classNames?: {
    button?: string;
  };
};

export const AddressItemList = ({
  classNames,
  offsetScrollbars = true,
}: Props) => {
  const queryClient = useQueryClient();

  const deliveryAddress = useMapStore((s) => s.deliveryAddress);
  const setDeliveryAddress = useMapStore((s) => s.setDeliveryAddress);
  const setIsMapOpen = useMapStore((s) => s.setIsMapOpen);

  const { data, isFetched } = useQuery({
    queryFn: getDeliveryAddresses,
    queryKey: [getDeliveryAddresses.queryKey],
  });

  const {
    mutate,
    isPending,
    variables: deletionUid,
  } = useMutation({
    mutationFn: deleteDeliveryAddress,
    mutationKey: [deleteDeliveryAddress.queryKey],
    onSuccess: (_, uid) => {
      if (uid === deliveryAddress?.uid) {
        setDeliveryAddress();
      }

      queryClient.invalidateQueries({
        queryKey: [getDeliveryAddresses.queryKey],
      });
      showSuccessNotification("Адрес успешно удален!");
    },
    onError: (error: AxiosError<any>) => {
      showErrorNotification(error);
    },
  });

  const handleOpenMap: MouseEventHandler<HTMLButtonElement> = (e) => {
    setIsMapOpen(true);
    e.stopPropagation();
  };

  useEffect(() => {
    if (data && data.data.length && !deliveryAddress) {
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

      <ScrollArea h={250} offsetScrollbars={offsetScrollbars} scrollbars="y">
        <Stack gap={12}>
          {data?.data?.map((address) => (
            <AddressItem
              disabled={isPending && address.uid === deletionUid}
              onDelete={() => mutate(address.uid)}
              onSelect={() => setDeliveryAddress(address)}
              active={address.uid === deliveryAddress?.uid}
              key={address.uid}
            >
              {getAddress(address)}
            </AddressItem>
          ))}
        </Stack>
      </ScrollArea>
    </Stack>
  );
};
