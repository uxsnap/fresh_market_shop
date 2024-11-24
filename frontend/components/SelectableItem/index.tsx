import { Text, Group, Paper, CloseButton } from "@mantine/core";
import { Circle } from "../icons/Circle";
import cn from "classnames";

import styles from "./SelectableItem.module.css";
import { CircleOk } from "../icons/CircleOk";

type SelectableItemIconProps = {
  onClick?: () => void;
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
          <Icon onClick={onMapOpen} fill="var(--mantine-color-accent-0)" />

          <Text maw="70%" truncate="end" c="accent.0" fz={16} fw={500}>
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
            onClick={onDelete}
            c="accent.0"
            size="lg"
          />
        </Group>
      </Group>
    </Paper>
  );
};
