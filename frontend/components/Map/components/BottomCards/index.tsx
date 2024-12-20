import { Box, Button, Stack, Text } from "@mantine/core";
import cn from "classnames";

import styles from "./BottomCards.module.css";
import { Carousel } from "@mantine/carousel";
import { Address } from "@/types";
import { PropsWithChildren } from "react";
import { useMapStore } from "@/store/map";
import { useQuery } from "@tanstack/react-query";
import { getDeliveryAddresses } from "@/api/user/getDeliveryAddresses";

const BottomCard = ({
  city,
  children,
  active,
}: PropsWithChildren<{ city: string; active: boolean }>) => (
  <Box className={cn(styles.bottomCard, active && styles.active)}>
    <Stack gap={8}>
      <Text lh="26px" fz={22} fw="bold" c="accent.0">
        {city}
      </Text>
      <Text lh="19px" fz={16} fw={500} c="accent.1">
        {children}
      </Text>
    </Stack>
  </Box>
);

export const BottomCards = () => {
  const { data } = useQuery({
    queryFn: getDeliveryAddresses,
    queryKey: [getDeliveryAddresses.queryKey],
  });

  const deliveryAddress = useMapStore((s) => s.deliveryAddress);
  const setDeliveryAddress = useMapStore((s) => s.setDeliveryAddress);
  const setIsFieldsModalOpen = useMapStore((s) => s.setIsFieldsModalOpen);

  return (
    <Stack justify="flex-end" gap={12} p={16} className={styles.root}>
      <Carousel slideGap="sm" align="start" dragFree withControls={false}>
        {(data?.data ?? []).map((item) => (
          <Carousel.Slide
            onClick={() => setDeliveryAddress(item)}
            key={item.uid}
            flex="1 0 auto"
          >
            <BottomCard
              active={item.addressUid === deliveryAddress?.addressUid}
              city={item.cityName ?? ""}
            >
              {item.streetName}
            </BottomCard>
          </Carousel.Slide>
        ))}
      </Carousel>

      <Button
        onClick={() => setIsFieldsModalOpen(true)}
        h={48}
        fz={18}
        variant="accent"
      >
        Добавить новый адрес
      </Button>
    </Stack>
  );
};
