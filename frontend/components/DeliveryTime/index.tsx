"use client";

import { getDelivery } from "@/api/delivery/getDelivery";
import { useCartStore } from "@/store";
import { useAuthStore } from "@/store/auth";
import { useMapStore } from "@/store/map";
import { formatDuration } from "@/utils";
import { Stack, Text } from "@mantine/core";
import { useQuery } from "@tanstack/react-query";
import { useEffect } from "react";

export const DeliveryTime = () => {
  const logged = useAuthStore((s) => s.logged);
  const deliveryAddress = useMapStore((s) => s.deliveryAddress);
  const setDelivery = useCartStore((s) => s.setDelivery);

  const { data } = useQuery({
    queryKey: [getDelivery.queryKey, deliveryAddress],
    queryFn: () =>
      getDelivery({
        orderUid: "00000000-0000-0000-0000-000000000000",
        deliveryAddressUid: deliveryAddress!.uid,
      }),
    enabled: !!logged && !!deliveryAddress,
  });

  useEffect(() => {
    if (!data?.data) {
      return;
    }

    setDelivery(data.data);
  }, [data?.data]);

  if (!deliveryAddress || data?.data.time === undefined) {
    return null;
  }

  const time = formatDuration(data?.data.time);

  return (
    <Stack pos="relative" gap={0} h={38} visibleFrom="sm">
      <Text style={{ whiteSpace: "nowrap" }} fw={500} size="md" c="accent.0">
        Около {!time ? "5 минут" : time}
      </Text>

      <Text style={{ whiteSpace: "nowrap" }} fw={500} size="xs" c="accent.2">
        Доставка: {data?.data.price}р
      </Text>
    </Stack>
  );
};
