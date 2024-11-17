import { Box, Button, Stack, Text } from "@mantine/core";

import styles from "./BottomCards.module.css";
import { Carousel } from "@mantine/carousel";
import { Address } from "@/types";
import { PropsWithChildren } from "react";

type Props = {
  items: Address[];
};

const BottomCard = ({
  city,
  children,
}: PropsWithChildren<{ city: string }>) => (
  <Box className={styles.bottomCard}>
    <Stack gap={8}>
      <Text fz={22} fw="bold" c="accent.0">
        {city}
      </Text>
      <Text fz={16} fw={500} c="accent.0">
        {children}
      </Text>
    </Stack>
  </Box>
);

export const BottomCards = ({ items }: Props) => (
  <Stack gap={12} p={16} className={styles.root}>
    <Carousel slideGap="sm" align="start" dragFree withControls={false}>
      {items.map((item) => (
        <Carousel.Slide key={item.uid} flex="1 0 auto">
          <BottomCard city={item.cityUid}>{item.street}</BottomCard>
        </Carousel.Slide>
      ))}
    </Carousel>

    <Button type="submit" h={48} fz={18} variant="accent">
      Добавить новый адрес
    </Button>
  </Stack>
);
