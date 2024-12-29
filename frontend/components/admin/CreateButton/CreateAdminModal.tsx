import { createAdminUser } from "@/api/auth/createAdmin";
import { getAdmins } from "@/api/auth/getAdmins";
import { showErrorNotification } from "@/utils";
import { Button, Group, PasswordInput, Stack, TextInput } from "@mantine/core";
import { useForm } from "@mantine/form";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { AxiosError } from "axios";

type Props = {
  onClose: () => void;
};

export const CreateAdminModal = ({ onClose }: Props) => {
  const queryClient = useQueryClient();

  const form = useForm({
    mode: "uncontrolled",
    initialValues: {
      name: "",
      email: "",
      password: "",
    },
    validate: {
      name: (value) =>
        value.trim().length >= 1
          ? null
          : "Длина имени пользователя должна быть больше 1",
      email: (value) =>
        /^\S+@\S+$/.test(value) ? null : "Неправильный формат email",
    },
  });

  const mutation = useMutation({
    mutationFn: createAdminUser,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [getAdmins.queryKey] });
      onClose();
    },
    onError: (error: AxiosError<any>) => {
      showErrorNotification(error);
    },
  });

  const handleSubmit = form.onSubmit((values) => {
    mutation.mutate(values);
  });

  return (
    <form onSubmit={handleSubmit}>
      <Stack gap={16}>
        <TextInput
          size="md"
          label="Имя и фамилия"
          placeholder="Введите имя и фамилию"
          {...form.getInputProps("name")}
        />

        <TextInput
          size="md"
          label="Email"
          placeholder="Введите email"
          {...form.getInputProps("email")}
        />
        <PasswordInput
          size="md"
          label="Пароль"
          placeholder="Введите пароль"
          {...form.getInputProps("password")}
        />
        <Group wrap="nowrap" mt={4} justify="space-between">
          <Button w="100%" type="submit" variant="accent">
            Создать
          </Button>
          <Button w="100%" onClick={() => onClose()} variant="accent-reverse">
            Закрыть
          </Button>
        </Group>
      </Stack>
    </form>
  );
};