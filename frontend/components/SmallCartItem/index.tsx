import { Group, Text, Image } from "@mantine/core";

import styles from "./SmallCartItem.module.css";
import { PropsWithChildren } from "react";

export const SmallCartItem = ({ children }: PropsWithChildren) => (
  <Group className={styles.root} px={8} align="center" gap={8}>
    <Image radius={8} mah={30} src="/recipe.png" />

    <Text fw={600} fz={18} c="accent.0" lh={1}>
      {children}
    </Text>
  </Group>
);
