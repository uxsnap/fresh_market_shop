import { Group, Text, Image } from "@mantine/core";

import styles from "./SmallCartItem.module.css";
import { getFallbackImg } from "@/utils";

type Props = {
  img?: string;
  children: string;
};

export const SmallCartItem = ({ children, img }: Props) => {
  const fallbackSrc = getFallbackImg(children);

  return (
    <Group className={styles.root} px={8} align="center" gap={8}>
      <Image radius={8} mah={30} src={img ?? fallbackSrc} />

      <Text fw={600} fz={18} c="accent.0" lh={1}>
        {children}
      </Text>
    </Group>
  );
};
