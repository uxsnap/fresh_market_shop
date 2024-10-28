import { AuthType } from "@/types";
import { Button, Flex } from "@mantine/core";

const mapTypeToActionButton = {
  login: "Войти",
  reg: "Зарегистрироваться",
  forgotPass: "Отправить",
  passRet: "Сохранить",
};

type Props = {
  currentType: AuthType;
  close: () => void;
};

export const Buttons = ({ currentType, close }: Props) => {
  const actionButton = mapTypeToActionButton[currentType];

  return (
    <Flex mt={16} gap={12} w="100%">
      <Button type="submit" fullWidth variant="accent">
        {actionButton}
      </Button>

      <Button onClick={close} fullWidth variant="secondary">
        Закрыть
      </Button>
    </Flex>
  );
};
