import cn from "classnames";
import { Stack, Text } from "@mantine/core";
import { Ok } from "../icons/Ok";
import { SadFace } from "../icons/SadFace";

import styles from "./HugeIconText.module.css";

type Props = {
  type: "ok" | "sad";
  children?: string;
  subText?: string;
  center?: boolean;
};

const mapTypeToIcon = {
  sad: <SadFace fill="var(--mantine-color-danger-0)" size={200} />,
  ok: <Ok fill="var(--mantine-color-primary-0)" size={200} />,
};

export const HugeIconText = ({ type, children, subText, center }: Props) => {
  const Icon = mapTypeToIcon[type];

  return (
    <Stack className={cn(center && styles.center)} gap={8} align="center">
      <div className={styles.icon}>{Icon}</div>

      <Text className={styles.text} c="accent.0" fw="bold" fz={32}>
        {children}
      </Text>

      {subText && (
        <Text className={styles.text} fw="bold" c="accent.2" fz={22}>
          {subText}
        </Text>
      )}
    </Stack>
  );
};
