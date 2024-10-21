"use client";

import { Card, Image, Text, Button, Container, Stack } from "@mantine/core";
import { Counter } from "../Counter";
import { ProductItem } from "@/types";
import { Carousel } from "@mantine/carousel";
import { getFallbackImg } from "@/utils";

type Props = ProductItem & {
  type?: "default" | "small";
  count: number;
  onIncrement: () => void;
  onDecrement: () => void;
  onAddItem: () => void;
};

const mapTypeToValues = {
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

export const ItemCard = ({
  type = "default",
  price,
  name,
  imgs = [],
  info,
  onIncrement,
  onDecrement,
  onAddItem,
  count,
}: Props) => {
  const { maw, imgH, priceFz, priceLh, infoFz, infoLh, nameFz, nameLh } =
    mapTypeToValues[type];

  const fallbackSrc = getFallbackImg(name);

  return (
    <Card p={8} w={maw} radius="md" withBorder>
      <Card.Section>
        {!imgs.length ? (
          <Image
            style={{ userSelect: "none" }}
            src={""}
            fallbackSrc={fallbackSrc}
            height={imgH}
            alt="Norway"
          />
        ) : (
          <Carousel withControls={false}>
            {imgs.map((img) => (
              <Carousel.Slide key={img}>
                <Image
                  style={{ userSelect: "none" }}
                  loading="lazy"
                  src={img}
                  height={imgH}
                  alt="Norway"
                  fit="contain"
                  fallbackSrc={fallbackSrc}
                />
              </Carousel.Slide>
            ))}
          </Carousel>
        )}
      </Card.Section>

      <Stack mt={8} gap={4}>
        <Text lh={`${priceLh}px`} fw={700} fz={priceFz} c="accent.0">
          {price} Руб.
        </Text>
        <Text lh={`${infoLh}px`} fw={500} fz={infoFz} c="accent.2">
          {info}
        </Text>
      </Stack>

      <Text lh={`${nameLh}px`} fw={500} fz={nameFz} mt={8} c="accent.0">
        {name}
      </Text>

      <Container fluid p={0} m={0} mt={8}>
        {count === 0 ? (
          <Button w="100%" onClick={onAddItem} variant="accent">
            Добавить
          </Button>
        ) : (
          <Counter
            count={count}
            onDecrement={onDecrement}
            onIncrement={onIncrement}
          />
        )}
      </Container>
    </Card>
  );
};
