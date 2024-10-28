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
import { useState } from "react";
import { Register } from "./components/Register";
import { Login } from "./components/Login";

const mapTypeToTitle = {
  login: "Вход в систему",
  reg: "Регистрация",
  forgotPass: "Забыли пароль",
  passRet: "Восстановление пароля",
};

const mapTypeToComponent = {
  login: (onChange: (type: AuthType) => void, close: () => void) => (
    <Login onChange={onChange} close={close} />
  ),
  reg: (onChange: (type: AuthType) => void, close: () => void) => (
    <Register onChange={onChange} close={close} />
  ),
  forgotPass: (onChange: (type: AuthType) => void) => (
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

const mapTypToText = {
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

type Props = {
  opened: boolean;
  close: () => void;
};

export const Auth = ({ opened, close }: Props) => {
  const [currentType, setCurrentType] = useState<AuthType>("login");

  const handleTypeChange = (type: AuthType) => {
    setCurrentType(type);
  };

  const title = mapTypeToTitle[currentType];
  const Component = mapTypeToComponent[currentType](handleTypeChange, close);
  const text = mapTypToText[currentType];

  return (
    <Modal.Root centered opened={opened} onClose={close}>
      <Modal.Overlay bg="accent.0" opacity={0.6} />

      <Modal.Content right={0}>
        <Modal.Body p={12}>
          <Title order={4} mb={12} c="accent.0">
            {title}
          </Title>

          {text && (
            <Text fz={12} lh="14px" c="accent.0">
              {text}
            </Text>
          )}

          {Component}
        </Modal.Body>
      </Modal.Content>
    </Modal.Root>
  );
};
