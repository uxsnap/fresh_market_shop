"use client";

import { getDelivery } from "@/api/delivery/getDelivery";
import { useAuthStore } from "@/store/auth";
import { useMapStore } from "@/store/map";
import { LoadingOverlay, Stack, Text } from "@mantine/core";
import { useQuery } from "@tanstack/react-query";

export const DeliveryTime = () => {
  const logged = useAuthStore((s) => s.logged);
  const deliveryAddress = useMapStore((s) => s.deliveryAddress);

  const { data, isFetching } = useQuery({
    queryKey: [getDelivery.queryKey],
    queryFn: () =>
      getDelivery({
        orderUid: "",
        deliveryAddressUid: deliveryAddress!.addressUid,
      }),
    enabled: !!logged && !!deliveryAddress,
  });

  if (!deliveryAddress) {
    return null;
  }

  return (
    <Stack gap={0} h={38} visibleFrom="sm">
      <LoadingOverlay
        visible={isFetching}
        zIndex={1}
        overlayProps={{ radius: "sm", blur: 2 }}
        loaderProps={{ color: "primary.0", type: "bars" }}
      />

      <Text style={{ whiteSpace: "nowrap" }} fw={500} size="md" c="accent.0">
        {data?.data.time}
      </Text>

      <Text style={{ whiteSpace: "nowrap" }} fw={500} size="xs" c="accent.2">
        Доставка: {data?.data.price}р
      </Text>
    </Stack>
  );
};
