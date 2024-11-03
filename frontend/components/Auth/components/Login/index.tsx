import { AuthType } from "@/types";
import { Button, Flex, Group, PasswordInput, TextInput } from "@mantine/core";
import { useForm } from "@mantine/form";
import { Buttons } from "../Buttons";
import { loginUser } from "@/api/auth/login";
import { useMutation } from "@tanstack/react-query";
import { useAuthStore } from "@/store/auth";

type Props = {
  onChange: (type: AuthType) => void;
  close: () => void;
};

export const Login = ({ onChange, close }: Props) => {
  const setLogged = useAuthStore((s) => s.setLogged);

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

  const mutation = useMutation({
    mutationFn: loginUser,
    onSuccess: () => {
      close();
      setLogged(true);
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
        <PasswordInput
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
          fz={10}
          h={12}
          size="xs"
          variant="outline"
        >
          Забыли пароль?
        </Button>
        <Button
          onClick={() => onChange("reg")}
          p={0}
          fz={10}
          h={12}
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
