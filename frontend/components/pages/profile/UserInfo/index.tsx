"use client";

import { getUserInfo } from "@/api/user/getUserInfo";
import { updateUser } from "@/api/user/updateUser";
import { Avatar } from "@/components/Avatar";
import { DateInput } from "@/components/DateInput";
import { ShadowBox } from "@/components/ShadowBox";
import { showSuccessNotification } from "@/utils";
import { Box, Button, Group, Stack, TextInput } from "@mantine/core";
import { useForm } from "@mantine/form";
import { useMutation, useQuery } from "@tanstack/react-query";
import { useEffect } from "react";

export const UserInfo = () => {
  const { data } = useQuery({
    queryFn: getUserInfo,
    queryKey: [getUserInfo.queryKey],
  });

  const form = useForm({
    mode: "uncontrolled",
    initialValues: {
      email: data?.data.email ?? "",
      firstName: "",
      lastName: "",
      birthday: "",
    },
    validate: {
      email: (value) =>
        /^\S+@\S+$/.test(value) ? null : "Неправильный формат email",
      firstName: (value) =>
        value.length > 0 ? null : "Длина имени должна быть больше 0",
      lastName: (value) =>
        value.length > 0 ? null : "Длина фамилии должна быть больше 0",
    },
  });

  useEffect(() => {
    form.setFieldValue("email", data?.data.email ?? "");
  }, [data]);

  const mutation = useMutation({
    mutationFn: updateUser,
    onSuccess: () => {
      showSuccessNotification("Пользователь успешно обновлен!");
    },
  });

  const handleSubmit = form.onSubmit((values) => {
    mutation.mutate(values);
  });

  return (
    <form onSubmit={handleSubmit}>
      <Box>
        <ShadowBox>
          <Stack gap={12}>
            <Box px={68} pt={20}>
              <Avatar />
            </Box>

            <Stack p={20} gap={12}>
              <TextInput
                size="md"
                label="Email"
                placeholder="Введите email"
                {...form.getInputProps("email")}
              />

              <TextInput
                size="md"
                label="Имя"
                placeholder="Введите имя"
                {...form.getInputProps("firstName")}
              />

              <TextInput
                size="md"
                label="Фамилия"
                placeholder="Введите фамилию"
                {...form.getInputProps("lastName")}
              />

              <DateInput {...form.getInputProps("birthday")} />
            </Stack>
          </Stack>
        </ShadowBox>

        <Group justify="space-between">
          <Button p={0} variant="outline">
            Выйти из системы
          </Button>
          <Button p={0} variant="outline" c="danger.0">
            Удалить аккаунт
          </Button>
        </Group>
      </Box>
    </form>
  );
};
