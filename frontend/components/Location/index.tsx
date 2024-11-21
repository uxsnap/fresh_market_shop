"use client";

import { Text, Group, Popover, useMatches } from "@mantine/core";
import { Location as LocationIcon } from "../icons/Location";

import styles from "./Location.module.css";
import { useCallback, useEffect, useState } from "react";
import { Map } from "../Map";
import { AddressItemList } from "../AddressItemList";

export const Location = () => {
  const [isMapOpen, setIsMapOpen] = useState(false);
  const [opened, setOpened] = useState(false);
  const [activeAddress, setActiveAddress] = useState("");

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

  return (
    <>
      <Popover
        opened={opened}
        width={500}
        position="bottom"
        withArrow
        shadow="md"
        closeOnClickOutside
      >
        <Popover.Target>
          <Group className={styles.group} onClick={handleOpen}>
            <LocationIcon />

            <Text truncate="end" fz={14} lh="150%" fw="bold" c="accent.0">
              {!activeAddress ? "Адрес не выбран" : activeAddress}
            </Text>
          </Group>
        </Popover.Target>

        <Popover.Dropdown>
          <AddressItemList
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
