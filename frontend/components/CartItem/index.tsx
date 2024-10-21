import { ProductItem } from "@/types";
import { getFallbackImg } from "@/utils";
import { Group, Paper, Image, Stack, Title, Text } from "@mantine/core";
import { Counter } from "../Counter";

type Props = ProductItem & {
  count: number;
  onDecrement: () => void;
  onIncrement: () => void;
};

export const CartItem = ({
  imgs,
  name,
  price,
  info,
  count,
  onDecrement,
  onIncrement,
}: Props) => {
  const fallbackSrc = getFallbackImg(name);

  return (
    <Paper radius="md" p={8}>
      <Group justify="space-between">
        <Group gap={16}>
          <Image
            loading="lazy"
            src={imgs[0]}
            height={80}
            alt={name}
            fit="contain"
            fallbackSrc={fallbackSrc}
          />

          <Stack gap={8}>
            <Title order={3} c="accent.0">
              {name}
            </Title>

            <Stack gap={4}>
              <Text c="accent.0">
                Цена: <Text fw="bold">{price}</Text>р
              </Text>

              <Text c="accent.2">
                Вес: <Text fw="bold">{info}</Text>р
              </Text>
            </Stack>
          </Stack>
        </Group>

        <Group gap={12}>
          <Counter
            onDecrement={onDecrement}
            onIncrement={onIncrement}
            count={count}
          />
        </Group>
      </Group>
    </Paper>
  );
};
