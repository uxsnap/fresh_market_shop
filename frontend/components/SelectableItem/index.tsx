import { Text, Group, Paper, CloseButton } from "@mantine/core";
import { Circle } from "../icons/Circle";
import cn from "classnames";

import styles from "./SelectableItem.module.css";
import { CircleOk } from "../icons/CircleOk";
import { MouseEventHandler } from "react";

type SelectableItemIconProps = {
  onClick?: (e: any) => void;
  fill: string;
};

type Props = {
  Icon: ({ onClick, fill }: SelectableItemIconProps) => JSX.Element;
  children?: string;
  onDelete?: () => void;
  onMapOpen?: () => void;
  onSelect?: () => void;
  active?: boolean;
  disabled?: boolean;
};

export const SelectableItem = ({
  children = "Очень длинный адрес доставки",
  onDelete,
  onMapOpen,
  active = false,
  disabled = false,
  onSelect,
  Icon,
}: Props) => {
  const handleDelete: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.stopPropagation();
    onDelete?.();
  };

  const handleMapOpen: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.stopPropagation();
    onMapOpen?.();
  };

  return (
    <Paper
      onClick={onSelect}
      className={cn(
        styles.root,
        active && styles.active,
        disabled && styles.disabled
      )}
      radius={8}
      withBorder
    >
      <Group
        wrap="nowrap"
        align="center"
        justify="space-between"
        py={8}
        px={12}
      >
        <Group wrap="nowrap" gap={16}>
          <Icon onClick={handleMapOpen} fill="var(--mantine-color-accent-0)" />

          <Text
            className={styles.text}
            truncate="end"
            c="accent.0"
            fz={16}
            fw={500}
          >
            {children}
          </Text>
        </Group>

        <Group w={62} wrap="nowrap" gap={8}>
          {active ? (
            <CircleOk size={20} />
          ) : (
            <Circle size={20} fill="var(--mantine-color-accent-0)" />
          )}

          <CloseButton
            p={0}
            iconSize={24}
            onClick={handleDelete}
            c="accent.0"
            size="lg"
          />
        </Group>
      </Group>
    </Paper>
  );
};
