import { Paper, PaperProps } from "@mantine/core";

import styles from "./ShadowBox.module.css";
import { PropsWithChildren } from "react";

export const ShadowBox = ({
  children,
  ...rest
}: PropsWithChildren<PaperProps>) => (
  <Paper radius={12} className={styles.root} {...rest}>
    {children}
  </Paper>
);
