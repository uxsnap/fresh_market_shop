import { Group, Paper, Text } from "@mantine/core";

import { Minus } from "../icons/Minus";
import { Plus } from "../icons/Plus";

import styles from "./Counter.module.css";

type Props = {
  onDecrement: () => void;
  onIncrement: () => void;
  count: number;
};

export const Counter = ({ count, onDecrement, onIncrement }: Props) => (
  <Paper w="100%" bg="primary.2" px={2}>
    <Group h={32} align="center" justify="space-between">
      <Paper
        display="flex"
        h={28}
        maw={28}
        style={{ cursor: "pointer" }}
        p={6}
        bg="bg.2"
        className={styles.button}
        onClick={onDecrement}
      >
        <Minus size={16} />
      </Paper>

      <Text className={styles.text} c="accent.0" fz={20} fw={500}>
        {count}
      </Text>

      <Paper
        display="flex"
        h={28}
        maw={28}
        style={{ cursor: "pointer" }}
        p={6}
        bg="bg.2"
        className={styles.button}
        onClick={onIncrement}
      >
        <Plus
          fill={`var(--mantine-color-${count === 10 ? "accent-2" : "accent-0"})`}
          size={16}
        />
      </Paper>
    </Group>
  </Paper>
);
