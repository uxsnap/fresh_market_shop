import { Stack, Text } from "@mantine/core";

type Props = {
  time?: string;
  price?: number;
};

export const DeliveryTime = ({ time = "1-2 часа", price = 180 }: Props) => {
  return (
    <Stack gap={0} h={38} visibleFrom="sm">
      <Text style={{ whiteSpace: "nowrap" }} fw={500} size="md" c="accent.0">
        {time}
      </Text>

      <Text style={{ whiteSpace: "nowrap" }} fw={500} size="xs" c="accent.2">
        Доставка: {price}р
      </Text>
    </Stack>
  );
};
