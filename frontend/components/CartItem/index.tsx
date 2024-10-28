import { getFallbackImg } from "@/utils";
import { Group, Paper, Image, Stack, Title, Text, Box } from "@mantine/core";
import { Counter } from "../Counter";
import { Trash } from "../icons/Trash";
import { memo } from "react";
import { useCartStore } from "@/store";
import { CartItem as ICartItem } from "@/types";

import styles from "./CartItem.module.css";

type Props = {
  item: ICartItem;
};

type CartItemInfoProps = {
  price: number;
  name: string;
  weight: number;
};

const CartItemCounter = memo(({ id, count }: { id: string; count: number }) => {
  const inc = useCartStore((state) => state.incCartItem);
  const dec = useCartStore((state) => state.decCartItem);
  const remove = useCartStore((state) => state.removeCartItem);

  return (
    <Group className={styles.counter} wrap="nowrap" gap={12} w="100%">
      <Counter
        onDecrement={() => dec(id)}
        onIncrement={() => inc(id)}
        count={count}
      />

      <Trash
        className={styles.icon}
        cursor="pointer"
        onClick={() => remove(id)}
        size={16}
        fill="var(--mantine-color-accent-0)"
      />
    </Group>
  );
});

const CartItemInfo = memo(({ name, price, weight }: CartItemInfoProps) => {
  return (
    <Stack gap={8}>
      <Title textWrap="balance" order={3} c="accent.0">
        {name}
      </Title>

      <Stack gap={4}>
        <Text fz={12} span c="accent.0">
          Цена:{" "}
          <Text fz={12} span fw="bold">
            {price}р
          </Text>
        </Text>

        <Text fz={12} span c="accent.2">
          Вес:{" "}
          <Text fz={12} span fw="bold">
            {weight}гр
          </Text>
        </Text>
      </Stack>
    </Stack>
  );
});

export const CartItem = memo(({ item }: Props) => {
  const { name, imgs, price, weight, id } = item.product;

  const fallbackSrc = getFallbackImg(name);

  const remove = useCartStore((state) => state.removeCartItem);

  return (
    <Paper radius="md" p={8}>
      <Group justify="space-between">
        <Group className={styles.mainGroup} gap={16} wrap="nowrap" justify="space-between">
          <Group w="100%" gap={16} wrap="nowrap">
            <Image
              loading="lazy"
              src={imgs[0]}
              height={80}
              w={80}
              alt={name}
              fit="contain"
              fallbackSrc={fallbackSrc}
              style={{ userSelect: "none" }}
            />

            <CartItemInfo price={price} name={name} weight={weight} />
          </Group>

          <Trash
            className={styles.mobileIcon}
            cursor="pointer"
            onClick={() => remove(id)}
            size={24}
            fill="var(--mantine-color-accent-0)"
          />
        </Group>

        <CartItemCounter id={id} count={item.count} />
      </Group>
    </Paper>
  );
});
