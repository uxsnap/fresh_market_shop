"use client";

import { Text, Group, Popover, useMatches, Box } from "@mantine/core";
import { Location as LocationIcon } from "../icons/Location";

import styles from "./Location.module.css";
import { useEffect, useState } from "react";
import { Map } from "../Map";
import { AddressItemList } from "../AddressItemList";
import { useClickOutside } from "@mantine/hooks";
import { getAddress } from "@/utils";
import { useMapStore } from "@/store/map";
import { useAuthStore } from "@/store/auth";
import { getDeliveryAddresses } from "@/api/user/getDeliveryAddresses";
import { useQueryClient } from "@tanstack/react-query";

export const Location = () => {
  const queryClient = useQueryClient();

  const [opened, setOpened] = useState(false);

  const logged = useAuthStore((s) => s.logged);
  const deliveryAddress = useMapStore((s) => s.deliveryAddress);

  // TODO: Разобраться почему не подтягивается state
  // const deliveryAddressesState = queryClient.getQueryState([
  //   getDeliveryAddresses.queryKey,
  // ]);

  const isMapOpen = useMapStore((s) => s.isMapOpen);
  const setIsMapOpen = useMapStore((s) => s.setIsMapOpen);

  const ref = useClickOutside(() => setOpened(false));

  const popupDisabled = useMatches({
    base: true,
    md: false,
  });

  const handleOpen = () => {
    setOpened(!popupDisabled);
    setIsMapOpen(popupDisabled);
  };

  useEffect(() => {
    if (isMapOpen) {
      setOpened(false);
    }
  }, [isMapOpen]);

  useEffect(() => {
    setOpened(false);
    setIsMapOpen(false);
  }, [popupDisabled]);

  // useEffect(() => {
  //   console.log(deliveryAddressesState);
  // }, [deliveryAddressesState]);

  if (!logged) {
    return null;
  }

  return (
    <>
      <Popover
        opened={opened}
        width={600}
        position="bottom"
        withArrow
        shadow="md"
        keepMounted
      >
        <Popover.Target>
          <Box w="100%">
            {logged && (
              <Group
                wrap="nowrap"
                ref={ref}
                className={styles.group}
                onClick={handleOpen}
              >
                <LocationIcon />

                <Text truncate="end" fz={14} lh="150%" fw="bold" c="accent.0">
                  {!deliveryAddress
                    ? "Адрес не выбран"
                    : getAddress(deliveryAddress)}
                </Text>
              </Group>
            )}
          </Box>
        </Popover.Target>

        <Popover.Dropdown pr={4}>
          <AddressItemList classNames={{ button: styles.addressButton }} />
        </Popover.Dropdown>
      </Popover>

      <Map />
    </>
  );
};
