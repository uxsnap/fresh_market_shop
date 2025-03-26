"use client";

import { Box, Flex, LoadingOverlay, Title } from "@mantine/core";
import { ItemCard } from "../ItemCard";
import { ProductItem } from "@/types";
import { Children, memo, PropsWithChildren } from "react";
import styles from "./ItemList.module.css";
import { useProductStore } from "@/store/product";
import { Carousel } from "@mantine/carousel";

type Props = {
  title?: string;
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
      <Carousel
        slideGap="sm"
        align="start"
        dragFree
        withControls={false}
        containScroll="trimSnaps"
      >
        {Children.map(children, (child) => (
          <Carousel.Slide flex="0 0 auto">{child}</Carousel.Slide>
        ))}
      </Carousel>
    );
  }

  return (
    <Box miw="100%">
      <Flex
        className={styles.wrapper}
        wrap={scroll ? "nowrap" : "wrap"}
        gap={12}
        align="flex-start"
      >
        {children}
      </Flex>
    </Box>
  );
};

export const ItemList = memo(
  ({
    title = "Вы уже заказывали",
    noTitle = false,
    items,
    isFetching = false,
    scroll = true,
  }: Props) => {
    const setCurItem = useProductStore((s) => s.setCurItem);

    return (
      <>
        <Flex
          className={styles.root}
          gap={20}
          pos="relative"
          direction="column"
        >
          {!noTitle && (
            <Title c="accent.0" order={1}>
              {title}
            </Title>
          )}

          <LoadingOverlay
            visible={isFetching}
            zIndex={1}
            overlayProps={{ radius: "sm", blur: 2 }}
            loaderProps={{ color: "primary.0", type: "bars" }}
          />

          {!isFetching && (
            <Wrapper scroll={scroll}>
              {(items ?? []).map((item, ind) => (
                <ItemCard
                  item={item}
                  key={ind}
                  onExtended={() => setCurItem(item)}
                />
              ))}
            </Wrapper>
          )}
        </Flex>
      </>
    );
  }
);
