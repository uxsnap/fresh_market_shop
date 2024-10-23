import { Button, Stack, Title, Text } from "@mantine/core";
import { TextWithPrice } from "../TextWithPrice";

import styles from "./PaymentBlock.module.css";

type Props = {
  text?: string;
  onClick: () => void;
};

export const PaymentBlock = ({ text = "Оплатить", onClick }: Props) => (
  <Stack miw="100%" gap={0}>
    <Title order={2} c="accent.0">
      Итого
    </Title>

    <Text mb={8} fw="bold" fz={14} c="accent.2">
      Доставка 1-2 часа
    </Text>

    <Stack className={styles.top} pb={12} gap={12}>
      <TextWithPrice text="За доставку" />
      {/* <TextWithPrice text="За товары"/> */}
      <TextWithPrice text="Сервисный сбор" price={10} />
    </Stack>

    <Stack mt={16} gap={12}>
      <TextWithPrice text="Всего" type="lg" />

      <Button h={50} fz={22} onClick={onClick} variant="accent">
        {text}
      </Button>
    </Stack>
  </Stack>
);
