import { ComponentProps, PropsWithChildren } from "react";

import styles from "./PopOver.module.css";
import { Paper } from "@mantine/core";
import cn from "classnames";

export const PopOver = ({
  children,
  className,
  ...rest
}: PropsWithChildren<
  ComponentProps<typeof Paper> & { className?: string }
>) => (
  <Paper radius={8} className={cn(styles.root, className)} {...rest}>
    {children}
  </Paper>
);
