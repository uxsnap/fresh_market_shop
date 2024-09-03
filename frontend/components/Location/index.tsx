import { Text, Flex } from "@mantine/core";
import { Location as LocationIcon } from "../icons/Location";

import styles from "./Location.module.css";

export const Location = () => (
  <Flex
    bd="1px solid var(--mantine-color-accent-0)"
    gap="sm"
    px={12}
    h={40}
    // TODO: Поменять значения на не абсолютные
    py={8}
    align="center"
    className={styles.root}
  >
    <LocationIcon />

    <Text maw={320} truncate="end" size="sm" fw="bold" c="accent.0">
      Адрес не выбран Адрес не выбран Адрес не выбран Адрес не выбран
    </Text>
  </Flex>
);
