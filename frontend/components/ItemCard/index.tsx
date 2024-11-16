"use client";

import { Card, Text, Stack, useMatches, Image } from "@mantine/core";
import { ProductItem } from "@/types";
import { memo, useMemo } from "react";

import { ItemCardIcon } from "./ItemCardIcon";
import styles from "./ItemCard.module.css";
import { ItemCounter } from "./ItemCounter";
import { getFallbackImg } from "@/utils";

type Props = {
  item: ProductItem;
  onExtended: () => void;
};

const mapTypeToValues: Record<string, any> = {
  default: {
    maw: 200,
    imgH: 176,
    priceFz: 22,
    priceLh: 26,
    infoFz: 12,
    infoLh: 14,
    nameFz: 14,
    nameLh: 16,
  },
  small: {
    maw: 140,
    imgH: 100,
    priceFz: 18,
    priceLh: 18,
    infoFz: 8,
    infoLh: 8,
    nameFz: 12,
    nameLh: 14,
  },
};

export const ItemCard = memo(({ item, onExtended }: Props) => {
  const type = useMatches({
    base: "small",
    md: "default",
  });

  const { maw, priceFz, priceLh, infoFz, infoLh, nameFz, nameLh } =
    mapTypeToValues[type];

  const { price, name, imgs = [], weight, ccal } = item;

  const fallbackSrc = useMemo(() => getFallbackImg(name), [name]);

  return (
    <Card p={8} w={maw} radius="md" withBorder pos="relative">
      <Card.Section>
        <ItemCardIcon type="max" onClick={onExtended} />

        <Image
          style={{ userSelect: "none" }}
          loading="lazy"
          src={imgs[0]}
          className={styles.img}
          alt={name}
          fit="contain"
          fallbackSrc={fallbackSrc}
          w="100%"
        />
      </Card.Section>

      <Stack mt={8} gap={4}>
        <Text lh={`${priceLh}px`} fw={700} fz={priceFz} c="accent.0">
          {price} Руб.
        </Text>
        <Text
          truncate="end"
          lh={`${infoLh}px`}
          fw={500}
          fz={infoFz}
          c="accent.2"
        >
          {weight} грамм/{ccal} ккал.
        </Text>
      </Stack>

      <Text
        truncate="end"
        lh={`${nameLh}px`}
        fw={500}
        fz={nameFz}
        mt={8}
        c="accent.0"
      >
        {name}
      </Text>

      <ItemCounter item={item} />
    </Card>
  );
});
