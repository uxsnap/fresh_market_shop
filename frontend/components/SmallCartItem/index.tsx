import { Group, Text, Image } from "@mantine/core";

import styles from "./SmallCartItem.module.css";
import { getFallbackImg } from "@/utils";

type Props = {
  img?: string;
  children: string;
  onClick: () => void;
};

export const SmallCartItem = ({ children, img, onClick }: Props) => {
  const fallbackSrc = getFallbackImg(children);

  const handleClick = () => {
    onClick();
  };

  return (
    <Group
      onClick={handleClick}
      className={styles.root}
      px={8}
      align="center"
      gap={8}
    >
      <Image
        radius={8}
        h={40}
        w={40}
        fit="contain"
        src={img ? `${process.env.NEXT_PUBLIC_API}/${img}` : fallbackSrc}
      />

      <Text fw={600} fz={18} c="accent.0" lh={1}>
        {children}
      </Text>
    </Group>
  );
};
