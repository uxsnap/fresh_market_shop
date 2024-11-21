"use client";

import { Text, Group, Popover, useMatches, Box } from "@mantine/core";
import { Location as LocationIcon } from "../icons/Location";

import styles from "./Location.module.css";
import { useCallback, useEffect, useState } from "react";
import { Map } from "../Map";
import { AddressItemList } from "../AddressItemList";
import { useClickOutside } from "@mantine/hooks";
import { useQuery } from "@tanstack/react-query";
import { getUserAddresses } from "@/api/user/getUserAdresses";
import { getAddress } from "@/utils";
import { UserAddress } from "@/types";

export const Location = () => {
  const [isMapOpen, setIsMapOpen] = useState(false);
  const [opened, setOpened] = useState(false);
  const [activeAddress, setActiveAddress] = useState<UserAddress>();

  const ref = useClickOutside(() => setOpened(false));

  const { data, isFetched } = useQuery({
    queryFn: getUserAddresses,
    queryKey: [getUserAddresses.queryKey],
  });

  const popupDisabled = useMatches({
    base: true,
    md: false,
  });

  const handleOpen = () => {
    setOpened(!popupDisabled);
    setIsMapOpen(popupDisabled);
  };

  const handleOpenMap = useCallback(() => {
    setIsMapOpen(true);
    setOpened(false);
  }, []);

  useEffect(() => {
    setOpened(false);
    setIsMapOpen(false);
  }, [popupDisabled]);

  useEffect(() => {
    if (data && data.data.length) {
      setActiveAddress(data.data[0]);
    }
  }, [isFetched]);

  return (
    <>
      <Popover
        opened={opened}
        width={500}
        position="bottom"
        withArrow
        shadow="md"
      >
        <Popover.Target>
          <Box w="100%">
            {isFetched && (
              <Group
                wrap="nowrap"
                ref={ref}
                className={styles.group}
                onClick={handleOpen}
              >
                <LocationIcon />

                <Text truncate="end" fz={14} lh="150%" fw="bold" c="accent.0">
                  {!activeAddress
                    ? "Адрес не выбран"
                    : getAddress(activeAddress)}
                </Text>
              </Group>
            )}
          </Box>
        </Popover.Target>

        <Popover.Dropdown pr={4}>
          <AddressItemList
            items={data?.data}
            onAdd={handleOpenMap}
            activeAddress={activeAddress}
            setActiveAddress={setActiveAddress}
          />
        </Popover.Dropdown>
      </Popover>

      <Map opened={isMapOpen} onClose={() => setIsMapOpen(false)} />
    </>
  );
};
