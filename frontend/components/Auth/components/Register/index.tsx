import { AuthType } from "@/types";
import { Button, Flex, Group, PasswordInput, TextInput } from "@mantine/core";
import { isEmail, useForm } from "@mantine/form";
import { Buttons } from "../Buttons";
import { useMutation } from "@tanstack/react-query";
import { registerUser } from "@/api/auth/register";
import { useRouter } from "next/navigation";
import { capitalize, showInlineErrorNotification } from "@/utils";
import { ChangeEvent } from "react";

type Props = {
  onChange: (type: AuthType) => void;
  close: () => void;
};

export const Register = ({ onChange, close }: Props) => {
  const router = useRouter();

  const form = useForm({
    mode: "uncontrolled",
    initialValues: {
      name: "",
      email: "",
      password: "",
    },
    validate: {
      name: (value) =>
        value.length >= 1
          ? null
          : "Длина имени пользователя должна быть больше 1",
      email: isEmail("Неправильный формат email"),
    },
  });

  const mutation = useMutation({
    mutationFn: registerUser,
    onSuccess: () => {
      close();
      router.push("/email_sent");
    },
    onError: () => {
      showInlineErrorNotification(
        "Ошибка регистрации, возможно такой email уже используется"
      );
    },
  });

  const handleSubmit = form.onSubmit((values) => {
    mutation.mutate(values);
  });

  const onNameChange = (e: ChangeEvent<HTMLInputElement>) => {
    const val = e.target.value.trimStart();

    if (!val.includes(" ")) {
      form.setFieldValue("name", capitalize(val));
      return;
    }

    const splitted = val.split(" ").filter((v) => !!v);

    form.setFieldValue(
      "name",
      `${capitalize(splitted[0])} ${capitalize(splitted[1])}`
    );
  };

  return (
    <form onSubmit={handleSubmit}>
      <Flex gap={16} direction="column">
        <TextInput
          size="md"
          label="Имя и Фамилия"
          placeholder="Введите имя и фамилию"
          {...form.getInputProps("name")}
          onChange={onNameChange}
          value={form.getValues().name}
        />

        <TextInput
          size="md"
          label="Email"
          placeholder="Введите email"
          {...form.getInputProps("email")}
        />
        <PasswordInput
          type="password"
          size="md"
          label="Пароль"
          placeholder="Введите пароль"
          {...form.getInputProps("password")}
        />
      </Flex>
      <Group mt={4} justify="space-between">
        <Button
          onClick={() => onChange("forgotPass")}
          p={0}
          fz={14}
          h={24}
          size="xs"
          variant="outline"
        >
          Забыли пароль?
        </Button>
        <Button
          onClick={() => onChange("login")}
          p={0}
          fz={14}
          h={24}
          size="xs"
          variant="outline"
        >
          Вход в систему
        </Button>
      </Group>

      <Buttons close={close} currentType="reg" />
    </form>
  );
};
