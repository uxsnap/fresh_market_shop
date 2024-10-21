"use client";

import {
  Container,
  Flex,
  LoadingOverlay,
  ScrollArea,
  Title,
} from "@mantine/core";
import { ItemCard } from "../ItemCard";
import { ProductItem } from "@/types";
import { PropsWithChildren } from "react";
import { useCartStore } from "@/store";

type Props = {
  title?: string;
  type?: "default" | "small";
  noTitle?: boolean;
  items?: ProductItem[];
  isFetching?: boolean;
  scroll?: boolean;
};

export const ItemList = ({
  type = "default",
  title = "Вы уже заказывали",
  noTitle = false,
  items = Array.from({ length: 10 }),
  isFetching = false,
  scroll = true,
}: Props) => {
  const { incCartItem, decCartItem, addCartItem, getCount } = useCartStore();

  const Wrapper = ({ children }: PropsWithChildren) => {
    if (scroll) {
      return (
        <ScrollArea type="never" w="100%">
          {children}
        </ScrollArea>
      );
    }

    return (
      <Container m={0} p={8}>
        {children}
      </Container>
    );
  };

  return (
    <Flex gap={20} mih={350} pos="relative" direction="column">
      {!noTitle && (
        <Title c="accent.0" order={1}>
          {title}
        </Title>
      )}

      <LoadingOverlay
        visible={isFetching}
        zIndex={1000}
        overlayProps={{ radius: "sm", blur: 2 }}
        loaderProps={{ color: "primary.0", type: "bars" }}
      />

      {!isFetching && (
        <Wrapper>
          <Flex wrap={scroll ? "nowrap" : "wrap"} gap={12} align="flex-start">
            {items.map((item, ind) => (
              <ItemCard
                {...item}
                type={type}
                key={ind}
                onAddItem={() => addCartItem(item)}
                onDecrement={() => decCartItem(item.id)}
                onIncrement={() => incCartItem(item.id)}
                count={getCount(item.id)}
              />
            ))}
          </Flex>
        </Wrapper>
      )}
    </Flex>
  );
};
