import { getFallbackImg } from "@/utils";
import { Group, Paper, Image, Stack, Title, Text } from "@mantine/core";
import { Counter } from "../Counter";
import { Trash } from "../icons/Trash";
import { memo } from "react";
import { useCartStore } from "@/store";
import { CartItem as ICartItem } from "@/types";

type Props = {
  item: ICartItem;
};

const CartItemCounter = memo(({ id, count }: { id: string; count: number }) => {
  const inc = useCartStore((state) => state.incCartItem);
  const dec = useCartStore((state) => state.decCartItem);
  const remove = useCartStore((state) => state.removeCartItem);

  return (
    <Group wrap="nowrap" gap={12} w="100%" maw={167}>
      <Counter
        onDecrement={() => dec(id)}
        onIncrement={() => inc(id)}
        count={count}
      />

      <Trash
        cursor="pointer"
        onClick={() => remove(id)}
        size={16}
        fill="var(--mantine-color-accent-0)"
      />
    </Group>
  );
});

export const CartItem = memo(({ item }: Props) => {
  const { name, imgs, price, weight, id } = item.product;

  const fallbackSrc = getFallbackImg(name);

  return (
    <Paper radius="md" p={8}>
      <Group justify="space-between">
        <Group gap={16}>
          <Image
            loading="lazy"
            src={imgs[0]}
            height={80}
            width={80}
            alt={name}
            fit="contain"
            fallbackSrc={fallbackSrc}
            style={{ userSelect: "none" }}
          />

          <Stack gap={8}>
            <Title order={3} c="accent.0">
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
        </Group>

        <CartItemCounter id={id} count={item.count} />
      </Group>
    </Paper>
  );
});
