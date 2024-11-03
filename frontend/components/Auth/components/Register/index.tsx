import { AuthType } from "@/types";
import { Button, Flex, Group, PasswordInput, TextInput } from "@mantine/core";
import { useForm } from "@mantine/form";
import { Buttons } from "../Buttons";
import { useMutation } from "@tanstack/react-query";
import { registerUser } from "@/api/auth/register";
import { useRouter } from "next/navigation";

type Props = {
  onChange: (type: AuthType) => void;
  close: () => void;
};

export const Register = ({ onChange, close }: Props) => {
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

  const mutation = useMutation({
    mutationFn: registerUser,
    onSuccess: () => {
      close();
      router.push("/email_sent");
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
          fz={10}
          h={12}
          size="xs"
          variant="outline"
        >
          Забыли пароль?
        </Button>
        <Button
          onClick={() => onChange("login")}
          p={0}
          fz={10}
          h={12}
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
