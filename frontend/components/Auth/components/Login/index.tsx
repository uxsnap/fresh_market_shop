import { AuthType } from "@/types";
import { Button, Flex, Group, TextInput } from "@mantine/core";
import { useForm } from "@mantine/form";
import { Buttons } from "../Buttons";
import { loginUser } from "@/api/auth/login";
import { useMutation } from "@tanstack/react-query";
import { useAuthStore } from "@/store/auth";
import { verifyUser } from "@/api/auth/verify";
import { useRouter } from "next/navigation";
import { showInlineErrorNotification } from "@/utils";

type Props = {
  onChange: (type: AuthType) => void;
  close: () => void;
};

export const Login = ({ onChange, close }: Props) => {
  const setLogged = useAuthStore((s) => s.setLogged);
  const setAdmin = useAuthStore((s) => s.setAdmin);

  const router = useRouter();

  const form = useForm({
    mode: "uncontrolled",
    initialValues: {
      email: "",
      password: "",
    },
    validate: {
      email: (value) =>
        /^\S+@\S+$/.test(value) ? null : "Неправильный формат email",
    },
  });

  const { mutate: mutateVerify } = useMutation({
    mutationFn: verifyUser,
    onSuccess: ({ isValid, isAdmin }) => {
      if (!isValid) {
        router.push("/");
      }

      setAdmin(isAdmin);
    },
    onError: () => {
      showInlineErrorNotification(
        "Ошибка верификации, возможно такого пользователя не существует"
      );
    },
  });

  const mutation = useMutation({
    mutationFn: loginUser,
    onSuccess: () => {
      close();
      setLogged(true);
      mutateVerify();

      router.push("/");
    },
    onError: () => {
      showInlineErrorNotification(
        "Ошибка логина, возможно такого пользователя не существует"
      );
    },
  });

  const handleSubmit = form.onSubmit((values) => {
    mutation.mutate(values);
  });

  return (
    <form onSubmit={handleSubmit}>
      <Flex gap={16} direction="column">
        <TextInput
          size="md"
          label="Email"
          placeholder="Введите email"
          {...form.getInputProps("email")}
        />
        <TextInput
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
          onClick={() => onChange("reg")}
          p={0}
          fz={14}
          h={24}
          size="xs"
          variant="outline"
        >
          Либо зарегистрируйтесь
        </Button>
      </Group>

      <Buttons close={close} currentType="login" />
    </form>
  );
};
