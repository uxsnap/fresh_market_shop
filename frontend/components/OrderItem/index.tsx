import { Group, Stack, Title, Text, Image } from "@mantine/core";
import { Refresh } from "../icons/Refresh";

export type Props = {
  title?: string;
  text?: string;
  price?: number;
  imgs?: string[];
  onClick?: () => void;
};

export const OrderItem = ({
  text = "Доставлен 22.04.23 в 18:00",
  title = "Заказ #1312123",
  price = 2400,
  imgs = Array.from({ length: 4 }),
  onClick,
}: Props) => {
  return (
    <Stack w="100%" p={12} bg="bg.1">
      <Group w="100%" justify="space-between" align="flex-start">
        <Group gap={12}>
          <Group style={{ cursor: "pointer" }} onClick={onClick}>
            <Refresh fill="var(--mantine-color-accent-0)" />
          </Group>

          <Stack gap={0}>
            <Title order={4} c="accent.0">
              {title}
            </Title>

            <Text fz={12} c="accent.3">
              {text}
            </Text>
          </Stack>
        </Group>

        <Title order={4} c="accent.0">
          {price} Руб
        </Title>
      </Group>

      <Group gap={12}>
        {imgs.map((img) => (
          <Image radius={8} key={img} mah={60} src="/recipe.png" />
        ))}
      </Group>
    </Stack>
  );
};
