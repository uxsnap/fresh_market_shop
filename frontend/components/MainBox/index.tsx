import { Box, BoxComponentProps } from "@mantine/core";
import { PropsWithChildren } from "react";

const SHELL_PADDING = 10;

export const MainBox = (props: PropsWithChildren<BoxComponentProps>) => {
  return (
    <Box
      pos="relative"
      h={`calc(100vh - var(--app-shell-header-height, 0px) - var(--app-shell-footer-height, 0px) - ${
        props.pt ?? 0
      }px - ${SHELL_PADDING}px)`}
      {...props}
    >
      {props.children}
    </Box>
  );
};
