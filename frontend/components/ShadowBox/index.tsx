import { Paper } from "@mantine/core";

import styles from "./ShadowBox.module.css";
import { PropsWithChildren } from "react";

export const ShadowBox = ({ children }: PropsWithChildren) => (
  <Paper radius={12} className={styles.root}>
    {children}
  </Paper>
);
