import { Group, Stack, Text } from "@mantine/core";
import styles from "./SuportItem.module.css";
import { Message } from "../icons/Message";
import { ISupportItem } from "@/types";
import { Closed } from "../icons/Closed";
import { Hourglass } from "../icons/Hourglass";
import { Ok } from "../icons/Ok";
import { dayJs } from "@/utils";

type Props = {
  name: string;
  description: string;
  date: string;
  status: ISupportItem["status"];
};

const mapStatusToText: Record<ISupportItem["status"], string> = {
  cant_solve: "Не выполнено",
  created: "В обработке",
  in_process: "В обработке",
  solved: "Выполнено",
};

const mapStatusToIcon: Record<ISupportItem["status"], JSX.Element> = {
  cant_solve: <Closed size={28} fill="var(--mantine-color-danger-2)" />,
  created: <Hourglass size={28} fill="var(--mantine-color-warning-0)" />,
  in_process: <Hourglass size={28} fill="var(--mantine-color-warning-0)" />,
  solved: <Ok size={28} fill="var(--mantine-color-primary-1)" />,
};

export const SupportItem = ({ name, description, date, status }: Props) => {
  const Icon = mapStatusToIcon[status];
  const text = mapStatusToText[status];

  return (
    <Group
      wrap="nowrap"
      align="center"
      justify="space-between"
      className={styles.root}
    >
      <Group wrap="nowrap" align="center" gap={16}>
        <Message size={32} fill="var(--mantine-color-accent-0)" />

        <Stack gap={8}>
          <Text lh={1} fw="bold" className={styles.name} c="accent.0">
            {name}
          </Text>
          <Text
            maw={200}
            truncate="end"
            lh={1}
            fw={600}
            className={styles.description}
            c="accent.1"
          >
            {description}
          </Text>

          <Text lh={1} fw={500} className={styles.date} c="accent.2">
            Дата обращения: {dayJs(date).format("DD.MM.YYYY")}
          </Text>
        </Stack>
      </Group>

      <Group wrap="nowrap" gap={8}>
        {Icon}

        <Text fz={14} c="accent.0" fw={600} className={styles.text}>
          {text}
        </Text>
      </Group>
    </Group>
  );
};
