import { Text, Group } from "@mantine/core";
import { Location as LocationIcon } from "../icons/Location";

import styles from "./Location.module.css";
import { useState } from "react";
import { Map } from "../Map";

export const Location = () => {
  const [isMapOpen, setIsMapOpen] = useState(false);

  return (
    <>
      <Group className={styles.root} onClick={() => setIsMapOpen(true)}>
        <LocationIcon />

        <Text truncate="end" fz={14} lh="150%" fw="bold" c="accent.0">
          Адрес не выбран
        </Text>
      </Group>

      <Map opened={isMapOpen} onClose={() => setIsMapOpen(false)} />
    </>
  );
};
