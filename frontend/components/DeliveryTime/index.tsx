import { Flex, Text } from "@mantine/core";

type Props = {
  time?: string 
  price?: number
};

export const DeliveryTime = ({ time = "1-2 часа", price = 180 }: Props) => {
  return (
    <Flex h={38} direction='column'>
      <Text fw={500} size="md" c="accent.0">
        {time}
      </Text>

      <Text fw={500} size="xs" c="accent.2">
        Доставка: {price}р
      </Text>
    </Flex>
  );
};