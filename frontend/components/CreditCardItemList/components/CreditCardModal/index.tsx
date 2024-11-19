import { Button, Group, Modal, Stack, TextInput, Title } from "@mantine/core";

import styles from "./CardCardModal.module.css";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { addPaymentCard } from "@/api/card/addPaymentCard";
import { showErrorNotification, showSuccessNotification } from "@/utils";
import { AxiosError } from "axios";
import { ErrorWrapper } from "@/types";
import { matches, useForm } from "@mantine/form";
import { getPaymentCardsByUser } from "@/api/card/getPaymentCardsByUser";

type Props = {
  opened: boolean;
  onClose: () => void;
};

export const CreditCardModal = ({ opened, onClose }: Props) => {
  const queryClient = useQueryClient();

  const form = useForm({
    mode: "controlled",
    initialValues: {
      number: "",
      expired: "",
      cvv: "",
    },
    onValuesChange: (values, previous) => {
      values.number = values.number.trim().replace(/[^0-9 ]/g, "");

      if ([4, 9, 14].includes(values.number.length)) {
        values.number += " ";
      }

      values.expired = values.expired.trim().replace(/[^0-9/]/g, "");

      if (values.expired.length === 2 && previous.expired.at(-1) !== "/") {
        values.expired += "/";
      }

      return values;
    },
    transformValues: (values) => ({
      ...values,
      number: values.number.replace(/ /g, ""),
    }),
    validate: {
      number: matches(/^\d{4} \d{4} \d{4} \d{4}$/, "Неправильный номер карты"),
      expired: matches(/^\d\d\/\d\d$/, "Неправильный срок окончания"),
      cvv: matches(/^[0-9]{3}$/, "Неправильный CVV"),
    },
  });

  const { mutate, isPending } = useMutation({
    mutationFn: addPaymentCard,
    onSuccess: () => {
      onClose();
      queryClient.invalidateQueries({
        queryKey: [getPaymentCardsByUser.queryKey],
      });
      showSuccessNotification("Карта успешно добавлена!");
    },
    onError: (error: AxiosError<{ error: ErrorWrapper }, any>) => {
      showErrorNotification(error);
    },
  });

  const handleSubmit = form.onSubmit((values) => {
    mutate(values);
  });

  const handleClose = () => {
    form.reset();
    onClose();
  };

  return (
    <Modal.Root centered opened={opened} onClose={handleClose}>
      <Modal.Overlay />

      <Modal.Content className={styles.content}>
        <Modal.Header className={styles.header}>
          <Title order={2} c="accent.0">
            Добавить карту
          </Title>
          <Modal.CloseButton size="32px" c="accent.0" />
        </Modal.Header>

        <Modal.Body px={20} py={12}>
          <form onSubmit={handleSubmit}>
            <Stack gap={12}>
              <TextInput
                required
                maxLength={19}
                size="md"
                placeholder="1234 5678 1234 5678"
                label="Номер карты"
                key={form.key("number")}
                {...form.getInputProps("number")}
              />

              <Group wrap="nowrap" w="100%" gap={12}>
                <TextInput
                  required
                  maxLength={5}
                  size="md"
                  label="Дата окончания"
                  placeholder="MM/ГГ"
                  w="100%"
                  key={form.key("expired")}
                  {...form.getInputProps("expired")}
                />
                <TextInput
                  required
                  maxLength={3}
                  size="md"
                  label="CVV"
                  placeholder="•••"
                  w="100%"
                  key={form.key("cvv")}
                  {...form.getInputProps("cvv")}
                />
              </Group>

              <Button
                variant="accent"
                disabled={isPending}
                w="100%"
                type="submit"
                mih={32}
              >
                Добавить
              </Button>
            </Stack>
          </form>
        </Modal.Body>
      </Modal.Content>
    </Modal.Root>
  );
};
