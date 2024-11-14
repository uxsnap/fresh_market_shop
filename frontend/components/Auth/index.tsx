import { AuthType } from "@/types";
import {
  Button,
  Flex,
  Group,
  Modal,
  TextInput,
  Title,
  Text,
} from "@mantine/core";
import { Register } from "./components/Register";
import { Login } from "./components/Login";
import { useAuthStore } from "@/store/auth";

const mapTypeToTitle: Record<AuthType, string> = {
  login: "Войдите для продолжения",
  reg: "Регистрация",
  forgotPass: "Забыли пароль",
  passRet: "Восстановление пароля",
};

const mapTypeToComponent: Record<
  AuthType,
  (onChange: (type: AuthType | "") => void) => any
> = {
  login: (onChange) => <Login onChange={onChange} close={() => onChange("")} />,
  reg: (onChange) => (
    <Register onChange={onChange} close={() => onChange("")} />
  ),
  forgotPass: (onChange) => (
    <>
      <Flex gap={16} direction="column">
        <TextInput size="md" label="Email" placeholder="Введите email" />
      </Flex>

      <Group mt={4} justify="space-between">
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
    </>
  ),
  passRet: () => (
    <>
      <Flex gap={16} direction="column">
        <TextInput size="md" label="Код" placeholder="Введите код" />
        <TextInput size="md" label="Пароль" placeholder="Введите пароль" />
      </Flex>
    </>
  ),
};

const mapTypToText: Record<AuthType, any> = {
  login: "",
  reg: "",
  forgotPass: (
    <>
      Введите email <br />
      На него будет отправлен код для восстановления
    </>
  ),
  passRet: "",
};

export const Auth = () => {
  const modalOpen = useAuthStore((s) => s.modalOpen);
  const setModalOpen = useAuthStore((s) => s.setModalOpen);

  const handleTypeChange = (type: AuthType | "") => {
    setModalOpen(type);
  };

  if (!modalOpen) {
    return null;
  }

  const title = mapTypeToTitle[modalOpen];
  const Component = mapTypeToComponent[modalOpen](handleTypeChange);
  const text = mapTypToText[modalOpen];

  return (
    <Modal.Root centered opened onClose={() => setModalOpen("")}>
      <Modal.Overlay bg="accent.0" opacity={0.6} />

      <Modal.Content right={0}>
        <Modal.Body p={12}>
          <Title order={2} mb={12} c="accent.0">
            {title}
          </Title>

          {text && (
            <Text fz={14} lh="14px" c="accent.0">
              {text}
            </Text>
          )}

          {Component}
        </Modal.Body>
      </Modal.Content>
    </Modal.Root>
  );
};
