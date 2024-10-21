import { ProductItem } from "@/types";
import { getFallbackImg } from "@/utils";
import { Group, Paper, Image, Stack, Title, Text } from "@mantine/core";
import { Counter } from "../Counter";
import { Trash } from "../icons/Trash";

type Props = ProductItem & {
  count: number;
  onDecrement: () => void;
  onIncrement: () => void;
  onDelete: () => void;
};

export const CartItem = ({
  imgs,
  name,
  price,
  weight,
  count,
  onDecrement,
  onIncrement,
  onDelete,
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
            width={80}
            alt={name}
            fit="contain"
            fallbackSrc={fallbackSrc}
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

        <Group wrap="nowrap" gap={12} w="100%" maw={167}>
          <Counter
            onDecrement={onDecrement}
            onIncrement={onIncrement}
            count={count}
          />

          <Trash
            cursor="pointer"
            onClick={onDelete}
            size={16}
            fill="var(--mantine-color-accent-0)"
          />
        </Group>
      </Group>
    </Paper>
  );
};
