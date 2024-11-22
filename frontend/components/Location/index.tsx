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

export const Location = () => {
  const [opened, setOpened] = useState(false);

  const logged = useAuthStore((s) => s.logged);
  const activeAddress = useMapStore((s) => s.activeAddress);

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

  if (!logged) {
    return null;
  }

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
            <Group
              wrap="nowrap"
              ref={ref}
              className={styles.group}
              onClick={handleOpen}
            >
              <LocationIcon />

              <Text truncate="end" fz={14} lh="150%" fw="bold" c="accent.0">
                {!activeAddress ? "Адрес не выбран" : getAddress(activeAddress)}
              </Text>
            </Group>
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
