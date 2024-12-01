"use client";

import { logoutUser } from "@/api/auth/logout";
import { getUser } from "@/api/user/getUser";
import { updateUser } from "@/api/user/updateUser";
import { Avatar } from "@/components/Avatar";
import { DateInput } from "@/components/DateInput";
import { ShadowBox } from "@/components/ShadowBox";
import { jwtError } from "@/constants";
import { useAuthStore } from "@/store/auth";
import { ErrorWrapper, User } from "@/types";
import { getErrorBody, isDateNull, showSuccessNotification } from "@/utils";
import { Box, Button, Group, Stack, TextInput } from "@mantine/core";
import { hasLength, isEmail, useForm } from "@mantine/form";
import { useMutation, useQuery } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { useCallback, useEffect } from "react";

import styles from "./UserInfo.module.css";
import { useRouter } from "next/navigation";
import { deleteAccount } from "@/api/auth/deleteAccount";
import { useCartStore } from "@/store";

type Form = {
  email: string;
  firstName: string;
  lastName: string;
  birthday: null | Date;
};

const getInitialValues = (data?: User) => ({
  email: data?.email ?? "",
  firstName: data?.firstName ?? "",
  lastName: data?.lastName ?? "",
  birthday: isDateNull(data?.birthday) ? null : new Date(data?.birthday + ""),
});

export const UserInfo = () => {
  const router = useRouter();
  const logged = useAuthStore((s) => s.logged);
  const setLogged = useAuthStore((s) => s.setLogged);
  const removeAllItems = useCartStore((s) => s.removeAllItems);

  const { data, error } = useQuery({
    queryFn: getUser,
    queryKey: [getUser.queryKey, logged],
    enabled: logged,
  });

  useEffect(() => {
    if (!error) {
      return;
    }

    const errorBody = getErrorBody(
      error as AxiosError<{ error: ErrorWrapper }>
    );

    if (errorBody?.type === jwtError) {
      logoutUser();
      removeAllItems();
    }
  }, [error]);

  const form = useForm<Form>({
    mode: "uncontrolled",
    initialValues: getInitialValues(data?.data),
    validate: {
      email: isEmail("Неправильный формат email"),
      firstName: hasLength({ min: 1 }, "Длина имени должна быть больше 1"),
    },
  });

  useEffect(() => {
    if (!data) {
      return;
    }

    form.initialize(getInitialValues(data?.data));
  }, [data]);

  const { mutate: logout, isPending: isPendingLogout } = useMutation({
    mutationFn: logoutUser,
    mutationKey: [logoutUser.queryKey],
    onSuccess: () => {
      router.push("/");
      setLogged(false);
    },
  });

  const { mutate: mutateDelete, isPending: isPendingDelete } = useMutation({
    mutationFn: deleteAccount,
    mutationKey: [deleteAccount.queryKey],
    onSuccess: () => {
      router.push("/");
      setLogged(false);
    },
  });

  const mutation = useMutation({
    mutationFn: updateUser,
    onSuccess: () => {
      form.resetDirty();
      showSuccessNotification("Пользователь успешно обновлен!");
    },
  });

  const handleSubmit = form.onSubmit((values) => {
    mutation.mutate(values as any);
  });

  const handleLogout = useCallback(() => {
    logout();
  }, [logout]);

  const handleDelete = useCallback(() => {
    mutateDelete();
  }, []);

  return (
    <form onSubmit={handleSubmit} className={styles.root}>
      <ShadowBox w="100%">
        <Stack className={styles.wrapper} gap={12}>
          <Box className={styles.avatarWrapper}>
            <Avatar upload />
          </Box>

          <Stack className={styles.inputs}>
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

            <DateInput clearable {...form.getInputProps("birthday")} />

            <Button
              disabled={!form.isDirty()}
              type="submit"
              w="100%"
              variant="accent"
            >
              Сохранить
            </Button>
          </Stack>
        </Stack>
      </ShadowBox>

      <Group justify="space-between">
        <Button
          disabled={isPendingLogout}
          onClick={handleLogout}
          p={0}
          variant="outline"
        >
          Выйти из системы
        </Button>
        <Button
          disabled={isPendingDelete}
          onClick={handleDelete}
          p={0}
          variant="outline"
          c="danger.0"
        >
          Удалить аккаунт
        </Button>
      </Group>
    </form>
  );
};
