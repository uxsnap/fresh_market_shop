import {
  Button,
  Flex,
  Group,
  Modal,
  TextInput,
  Title,
  Text,
} from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import { useState } from "react";

type AuthType = "login" | "reg" | "forgotPass" | "passRet";

const mapTypeToActionButton = {
  login: "Войти",
  reg: "Зарегистрироваться",
  forgotPass: "Отправить",
  passRet: "Сохранить",
};

const mapTypeToTitle = {
  login: "Вход в систему",
  reg: "Регистрация",
  forgotPass: "Забыли пароль",
  passRet: "Восстановление пароля",
};

const mapTypeToComponent = {
  login: (onChange: (type: AuthType) => void) => (
    <>
      <Flex gap={16} direction="column">
        <TextInput size="md" label="Email" placeholder="Введите email" />
        <TextInput size="md" label="Пароль" placeholder="Введите пароль" />
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
    </>
  ),
  reg: (onChange: (type: AuthType) => void) => (
    <>
      <Flex gap={16} direction="column">
        <TextInput size="md" label="Имя" placeholder="Введите имя" />
        <TextInput size="md" label="Email" placeholder="Введите email" />
        <TextInput size="md" label="Пароль" placeholder="Введите пароль" />
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
    </>
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

export const Auth = () => {
  const [opened, { open, close }] = useDisclosure(false);
  const [currentType, setCurrentType] = useState<AuthType>("login");

  const handleTypeChange = (type: AuthType) => {
    setCurrentType(type);
  };

  const actionButton = mapTypeToActionButton[currentType];
  const title = mapTypeToTitle[currentType];
  const Component = mapTypeToComponent[currentType](handleTypeChange);
  const text = mapTypToText[currentType];

  return (
    <Modal.Root opened={true} onClose={close}>
      <Modal.Overlay />

      <Modal.Content>
        <Modal.Body p={12}>
          <Title mb={12} c="accent.0">
            {title}
          </Title>

          {text && (
            <Text fz={12} lh="14px" c="accent.0">
              {text}
            </Text>
          )}

          {Component}

          <Flex mt={16} gap={12} w="100%">
            <Button fullWidth variant="accent">
              {actionButton}
            </Button>

            <Button fullWidth variant="secondary">
              Закрыть
            </Button>
          </Flex>
        </Modal.Body>
      </Modal.Content>
    </Modal.Root>
  );
};
