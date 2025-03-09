"use client";

import { getDelivery } from "@/api/delivery/getDelivery";
import { useCartStore } from "@/store";
import { useAuthStore } from "@/store/auth";
import { useMapStore } from "@/store/map";
import { dayJs } from "@/utils";
import { LoadingOverlay, Stack, Text } from "@mantine/core";
import { useQuery } from "@tanstack/react-query";
import { useEffect } from "react";

export const DeliveryTime = () => {
  const logged = useAuthStore((s) => s.logged);
  const deliveryAddress = useMapStore((s) => s.deliveryAddress);
  const setDelivery = useCartStore((s) => s.setDelivery);

  const { data, isFetching } = useQuery({
    queryKey: [getDelivery.queryKey],
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

  if (!deliveryAddress) {
    return null;
  }

  const time = dayJs(data?.data.time);
  const formatted = time.format("Hч mmм");

  return (
    <Stack gap={0} h={38} visibleFrom="sm">
      <LoadingOverlay
        visible={isFetching}
        zIndex={1}
        overlayProps={{ radius: "sm", blur: 2 }}
        loaderProps={{ color: "primary.0", type: "bars" }}
      />

      <Text style={{ whiteSpace: "nowrap" }} fw={500} size="md" c="accent.0">
        Около {formatted}
      </Text>

      <Text style={{ whiteSpace: "nowrap" }} fw={500} size="xs" c="accent.2">
        Доставка: {data?.data.price}р
      </Text>
    </Stack>
  );
};
