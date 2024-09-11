import { Flex, Title } from "@mantine/core";
import { ItemCard } from "../ItemCard";

type Props = {
  title?: string;
  type?: "default" | "small";
  noTitle?: boolean;
};

export const ItemList = ({
  type = "default",
  title = "Вы уже заказывали",
  noTitle = false,
}: Props) => {
  return (
    <Flex gap={20} direction="column">
      {!noTitle && <Title order={1}>{title}</Title>}

      <Flex style={{ overflowX: "auto" }} gap={12} align="flex-start">
        {Array.from({ length: 10 }).map((_, ind) => (
          <ItemCard type={type} key={ind} />
        ))}
      </Flex>
    </Flex>
  );
};
