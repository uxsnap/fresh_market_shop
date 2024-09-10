import { Button, Stack } from "@mantine/core";
import { TextWithPrice } from "../TextWithPrice";

import styles from "./PaymentBlock.module.css";

export const PaymentBlock = () => {
  return (
    <Stack>
      <Stack className={styles.top} pb={12} gap={12}>
        <TextWithPrice />
        <TextWithPrice />
        <TextWithPrice />
      </Stack>

      <Stack gap={12}>
        <TextWithPrice type="lg" />

        <Button variant="accent">Оплатить</Button>
      </Stack>
    </Stack>
  );
};
