import { Button, Stack, Title, Text } from "@mantine/core";
import { TextWithPrice } from "../TextWithPrice";

import styles from "./PaymentBlock.module.css";
import { useCartStore } from "@/store";
import { useEffect, useState } from "react";
import { formatDuration } from "@/utils";

type Props = {
  buttonText?: string;
  onClick?: () => void;
  disabled?: boolean;
  price?: number;
};

export const PaymentBlock = ({
  buttonText = "Оплатить",
  onClick,
  disabled = false,
  price,
}: Props) => {
  const storePrice = useCartStore((s) => s.getItemsPrice());
  const delivery = useCartStore((s) => s.delivery);

  const [curPrice, setCurPrice] = useState(0);

  useEffect(() => {
    setCurPrice(price !== undefined ? price : storePrice);
  }, [price, storePrice]);

  const calculateDelivery = () => {
    if (!delivery) {
      return null;
    }

    const time = formatDuration(delivery.time / 1000);

    return (
      <Text mb={8} fw="bold" fz={14} c="accent.2">
        Доставка около {!time ? "5 минут" : time}
      </Text>
    );
  };

  return (
    <Stack visibleFrom="md" miw="100%" gap={0}>
      <Title order={2} c="accent.0">
        Итого
      </Title>

      {calculateDelivery()}

      <Stack className={styles.top} pb={12} gap={12}>
        {delivery && (
          <TextWithPrice text="За доставку" price={delivery?.price} />
        )}
        <TextWithPrice text="За товары" price={curPrice} />
        <TextWithPrice text="Сервисный сбор" price={10} />
      </Stack>

      <Stack mt={16} gap={12}>
        <TextWithPrice
          text="Всего"
          price={curPrice + (delivery?.price ?? 0) + 10}
          type="lg"
        />

        <Button
          disabled={disabled}
          h={50}
          fz={22}
          onClick={onClick}
          variant="accent"
        >
          {buttonText}
        </Button>
      </Stack>
    </Stack>
  );
};
