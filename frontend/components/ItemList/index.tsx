import { Flex, Title } from "@mantine/core";
import { ItemCard } from "../ItemCard";

type Props = {
  title?: string;
};

export const ItemList = ({ title = "Вы уже заказывали" }: Props) => {
  return (
    <Flex gap={20} direction="column">
      <Title order={1}>{title}</Title>

      <Flex gap={12} align="flex-start">
        {Array.from({ length: 10 }).map((_, ind) => (
          <ItemCard key={ind} />
        ))}
      </Flex>
    </Flex>
  );
};
