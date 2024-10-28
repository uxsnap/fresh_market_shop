import { AuthType } from "@/types";
import { Button, Flex, Group, PasswordInput, TextInput } from "@mantine/core";
import { useForm } from "@mantine/form";
import { Buttons } from "../Buttons";

type Props = {
  onChange: (type: AuthType) => void;
  close: () => void;
};

export const Login = ({ onChange, close }: Props) => {
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

  return (
    <form onSubmit={form.onSubmit((values) => console.log(values))}>
      <Flex gap={16} direction="column">
        <TextInput size="md" label="Email" placeholder="Введите email" />
        <PasswordInput size="md" label="Пароль" placeholder="Введите пароль" />
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
          Регистрация
        </Button>
      </Group>

      <Buttons close={close} currentType="reg" />
    </form>
  );
};
