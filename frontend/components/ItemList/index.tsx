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
import { memo, PropsWithChildren } from "react";

type Props = {
  title?: string;
  type?: "default" | "small";
  noTitle?: boolean;
  items?: ProductItem[];
  isFetching?: boolean;
  scroll?: boolean;
};

const Wrapper = ({
  scroll,
  children,
}: PropsWithChildren<{ scroll: boolean }>) => {
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

export const ItemList = memo(
  ({
    type = "default",
    title = "Вы уже заказывали",
    noTitle = false,
    items = Array.from({ length: 10 }),
    isFetching = false,
    scroll = true,
  }: Props) => (
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
        <Wrapper scroll={scroll}>
          <Flex wrap={scroll ? "nowrap" : "wrap"} gap={12} align="flex-start">
            {items.map((item, ind) => (
              <ItemCard item={item} type={type} key={ind} />
            ))}
          </Flex>
        </Wrapper>
      )}
    </Flex>
  )
);
