import { Flex, LoadingOverlay, ScrollArea, Title } from "@mantine/core";
import { ItemCard } from "../ItemCard";
import { ProductItem } from "@/types";

type Props = {
  title?: string;
  type?: "default" | "small";
  noTitle?: boolean;
  items?: ProductItem[];
  isFetching?: boolean;
};

export const ItemList = ({
  type = "default",
  title = "Вы уже заказывали",
  noTitle = false,
  items = Array.from({ length: 10 }),
  isFetching = false,
}: Props) => {
  return (
    <Flex gap={20} mih={350} pos="relative" direction="column">
      {!noTitle && <Title order={1}>{title}</Title>}

      <LoadingOverlay
        visible={isFetching}
        zIndex={1000}
        overlayProps={{ radius: "sm", blur: 2 }}
        loaderProps={{ color: "primary.0", type: "bars" }}
      />

      {!isFetching && (
        <ScrollArea type="never" w="100%">
          <Flex gap={12} align="flex-start">
            {items.map((item, ind) => (
              <ItemCard {...item} type={type} key={ind} />
            ))}
          </Flex>
        </ScrollArea>
      )}
    </Flex>
  );
};
