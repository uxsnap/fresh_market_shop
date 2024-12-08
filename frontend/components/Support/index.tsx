import {
  Button,
  Group,
  Input,
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
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { getAllTopics } from "@/api/support/getAllTopics";
import { addTicket } from "@/api/support/addTicket";
import { IMaskInput } from "react-imask";
import { useState } from "react";
import { getTickets } from "@/api/support/getTickets";

export const Support = () => {
  const [opened, setOpened] = useState(false);

  const queryClient = useQueryClient();

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

  const { mutate, isPending } = useMutation({
    mutationFn: addTicket,
    mutationKey: [addTicket.queryKey],
    onSuccess: () => {
      setOpened(false);
      queryClient.invalidateQueries({
        queryKey: [getTickets.queryKey],
      });
    },
  });

  const form = useForm({
    mode: "uncontrolled",
    initialValues: {
      topicUid: "",
      title: "",
      description: "",
      fromEmail: "",
      fromPhone: "",
    },
    validate: {
      topicUid: isNotEmpty("Заполните тему обращения"),
      title: isNotEmpty("Заполните заголовок"),
      description: isNotEmpty("Заполните описание проблемы"),
    },
  });

  const handleSubmit = form.onSubmit((values) => {
    mutate(values);
  });

  return (
    <Popover
      trapFocus
      position="left"
      withArrow
      shadow="md"
      offset={{ mainAxis: 10, crossAxis: -100 }}
      opened={opened}
      onChange={setOpened}
    >
      <Popover.Target>
        <Group
          onClick={() => setOpened((o) => !o)}
          wrap="nowrap"
          gap={8}
          className={styles.root}
        >
          <Text span fz={14} fw="bold" c="accent.0">
            Написать нам
          </Text>
          <MailSent fill="var(--mantine-color-accent-0)" size={24} />
        </Group>
      </Popover.Target>

      <Popover.Dropdown>
        <form onSubmit={handleSubmit} className={styles.form}>
          <Title mb={16} order={3} c="accent.0">
            Свяжитесь с нами
          </Title>

          <Stack gap={12}>
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
              comboboxProps={{ withinPortal: false }}
            />

            <TextInput
              placeholder="Введите заголовок"
              withAsterisk
              key={form.key("title")}
              size="md"
              {...form.getInputProps("title")}
            />

            <TextInput
              size="md"
              withAsterisk
              placeholder="Введите email"
              {...form.getInputProps("fromEmail")}
            />

            <Input
              size="md"
              label="Телефон"
              component={IMaskInput}
              mask="+7 (000) 000-00-00"
              placeholder="Введите телефон"
              {...form.getInputProps("fromPhone")}
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
              disabled={isFetching || isPending}
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
