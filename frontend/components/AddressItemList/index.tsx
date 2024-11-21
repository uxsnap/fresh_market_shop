"use client";

import { Button, LoadingOverlay, Stack } from "@mantine/core";
import { AddressItem } from "../AddressItem";
import { Plus } from "../icons/Plus";
import { useQuery } from "@tanstack/react-query";
import { getUserAddresses } from "@/api/user/getUserAdresses";

type Props = {
  activeAddress: string;
  setActiveAddress: (v: string) => void;
  onAdd: () => void;
};

export const AddressItemList = ({
  activeAddress,
  setActiveAddress,
  onAdd,
}: Props) => {
  const { data, isFetching } = useQuery({
    queryFn: getUserAddresses,
    queryKey: [getUserAddresses.queryKey],
  });

  return (
    <Stack gap={12}>
      <Button
        onClick={onAdd}
        mih={48}
        variant="dashed"
        leftSection={<Plus fill="var(--mantine-color-accent-0)" />}
      >
        Добавить
      </Button>

      <LoadingOverlay
        visible={isFetching}
        zIndex={1}
        overlayProps={{ radius: "sm", blur: 2 }}
        loaderProps={{ color: "primary.0", type: "bars" }}
      />

      {(data?.data ?? []).map((address) => (
        <AddressItem
          onSelect={() => setActiveAddress(address.addressUid)}
          active={address.addressUid === activeAddress}
          key={address.addressUid}
        />
      ))}
    </Stack>
  );
};
