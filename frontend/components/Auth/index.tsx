import { AuthType } from "@/types";
import {
  Button,
  Flex,
  Group,
  Modal,
  TextInput,
  Title,
  Text,
  Box,
} from "@mantine/core";
import { Register } from "./components/Register";
import { Login } from "./components/Login";
import { useAuthStore } from "@/store/auth";
import { Buttons } from "./components/Buttons";
import { ForgotPass } from "./components/ForgotPass";

const mapTypeToTitle: Record<AuthType, string> = {
  login: "Войдите для продолжения",
  reg: "Регистрация",
  forgotPass: "Восстановление пароля",
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
    <ForgotPass onChange={onChange} close={() => onChange("")} />
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
    <Box mb={20}>
      Введите email, на него будет отправлен код для восстановления
    </Box>
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
