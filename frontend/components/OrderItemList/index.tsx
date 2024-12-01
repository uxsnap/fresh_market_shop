"use client";

import { Box, LoadingOverlay, Stack } from "@mantine/core";
import { OrderItem } from "../OrderItem";
import { useQuery } from "@tanstack/react-query";
import { getOrdersHistory } from "@/api/order/getOrderHistory";

export const OrderItemList = () => {
  const { data, isFetching } = useQuery({
    queryFn: getOrdersHistory,
    queryKey: [getOrdersHistory.queryKey],
  });

  console.log(data);

  return (
    <Box>
      <LoadingOverlay
        visible={isFetching}
        zIndex={1}
        overlayProps={{ radius: "sm", blur: 2 }}
        loaderProps={{ color: "primary.0", type: "bars" }}
      />

      <Stack gap={12}>
        {data?.data.map((order, ind) => <OrderItem key={ind} {...order} />)}
      </Stack>
    </Box>
  );
};
