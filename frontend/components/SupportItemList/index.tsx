"use client";

import { LoadingOverlay, Stack, Title } from "@mantine/core";
import { SupportItem } from "../SupportItem";
import { useQuery } from "@tanstack/react-query";
import { getTickets } from "@/api/support/getTickets";
import { memo } from "react";
import { ShadowBox } from "../ShadowBox";

export const SupportItemList = memo(() => {
  const { data, isFetching } = useQuery({
    queryFn: getTickets,
    queryKey: [getTickets.queryKey],
  });

  if (!data?.data.length) {
    return null;
  }

  return (
    <ShadowBox
      p={12}
      w="100%"
      mih={312}
      mah={312}
      style={{ overflowY: "auto" }}
      pos="relative"
    >
      <Stack gap={16}>
        <Title c="accent.0" order={3}>
          Обращения
        </Title>

        <LoadingOverlay
          visible={isFetching}
          zIndex={1}
          overlayProps={{ radius: "sm", blur: 2 }}
          loaderProps={{ color: "primary.0", type: "bars" }}
        />

        <Stack gap={12}>
          {data?.data.map((item) => (
            <SupportItem
              key={item.uid}
              name={item.title}
              description={item.description}
              date={item.createdAt}
              status={item.status}
            />
          ))}
        </Stack>
      </Stack>
    </ShadowBox>
  );
});
