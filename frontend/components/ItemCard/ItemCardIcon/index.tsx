import { ArrowsMinimize } from "@/components/icons/ArrowsMinimize";
import { ArrowsMaximize } from "@/components/icons/ArrowsMaximize";
import { memo } from "react";
import styles from "./ItemCardIcon.module.css";
import cn from "classnames";
import { Box } from "@mantine/core";

type Props = {
  type: "min" | "max";
  onClick: () => void;
};

export const ItemCardIcon = memo(({ type, onClick }: Props) => (
  <Box display="flex" onClick={onClick}>
    {type === "max" && (
      <ArrowsMaximize
        className={cn(styles.icon, styles.max)}
        fill="var(--mantine-color-accent-0)"
      />
    )}

    {type === "min" && (
      <ArrowsMinimize
        className={styles.icon}
        fill="var(--mantine-color-accent-0)"
      />
    )}
  </Box>
));
