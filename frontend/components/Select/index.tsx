"use client";

import cn from "classnames";
import { Select as MantineSelect } from "@mantine/core";
import { ComponentProps } from "react";
import { ArrowDown } from "../icons/ArrowDown";

import styles from "./Select.module.css";
import { useDisclosure } from "@mantine/hooks";

export const Select = (props: ComponentProps<typeof MantineSelect>) => {
  const [dropdownOpened, { toggle }] = useDisclosure();

  return (
    <MantineSelect
      leftSection={
        <div className={cn(styles.icon, dropdownOpened && styles.rotated)}>
          <ArrowDown size={20} fill="var(--mantine-color-accent-0)" />
        </div>
      }
      rightSection={<></>}
      withCheckIcon={false}
      classNames={{
        options: styles.options,
        option: styles.option,
      }}
      onClick={toggle}
      {...props}
    />
  );
};
