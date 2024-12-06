import {
  Button,
  Group,
  Popover,
  Select,
  Stack,
  Text,
  Textarea,
  TextInput,
  Title,
} from "@mantine/core";

import styles from "./Support.module.css";
import { MailSent } from "../icons/MailSent";
import { isNotEmpty, useForm } from "@mantine/form";
import { useQuery } from "@tanstack/react-query";
import { getAllTopics } from "@/api/support/getAllTopics";

export const Support = () => {
  const { data, isFetching } = useQuery({
    queryFn: getAllTopics,
    queryKey: [getAllTopics.queryKey],
    select(data) {
      return data.data.map((item) => ({
        label: item.name,
        value: item.uid,
      }));
    },
  });

  const form = useForm({
    mode: "uncontrolled",
    initialValues: {
      name: "",
      description: "",
    },
    validate: {
      name: isNotEmpty("Заполните название проблемы"),
      description: isNotEmpty("Заполните описание проблемы"),
    },
  });

  return (
    <Popover
      trapFocus
      position="left"
      withArrow
      shadow="md"
      offset={{ mainAxis: 10, crossAxis: -100 }}
    >
      <Popover.Target>
        <Group wrap="nowrap" gap={8} className={styles.root}>
          <Text span fz={14} fw="bold" c="accent.0">
            Написать нам
          </Text>
          <MailSent fill="var(--mantine-color-accent-0)" size={24} />
        </Group>
      </Popover.Target>

      <Popover.Dropdown>
        <form className={styles.form}>
          <Title mb={16} order={3} c="accent.0">
            Свяжитесь с нами
          </Title>

          <Stack gap={16}>
            <Select
              w="100%"
              size="md"
              label="Тема обращения"
              placeholder="Выберите тему обращения"
              data={data ?? []}
              allowDeselect={false}
              withAsterisk
              withScrollArea={false}
              styles={{ dropdown: { maxHeight: 130, overflowY: "auto" } }}
              key={form.key("topicUid")}
              {...form.getInputProps("topicUid")}
            />

            <Textarea
              radius="md"
              label="Описание проблемы"
              placeholder="Введите описание проблемы"
              {...form.getInputProps("description")}
              resize="vertical"
              minRows={10}
              classNames={{ wrapper: styles.textarea }}
            />

            <Button
              variant="accent"
              disabled={isFetching}
              w="100%"
              type="submit"
              mih={32}
            >
              Отправить
            </Button>
          </Stack>
        </form>
      </Popover.Dropdown>
    </Popover>
  );
};
