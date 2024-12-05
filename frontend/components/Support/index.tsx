import {
  Button,
  Group,
  Popover,
  Stack,
  Text,
  Textarea,
  TextInput,
  Title,
} from "@mantine/core";

import styles from "./Support.module.css";
import { MailSent } from "../icons/MailSent";
import { isNotEmpty, useForm } from "@mantine/form";

export const Support = () => {
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
            <TextInput
              size="md"
              label="Название проблемы"
              placeholder="Введите название проблемы"
              {...form.getInputProps("name")}
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
              // disabled={isPending}
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
