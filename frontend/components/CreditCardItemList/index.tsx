"use client";

import { Button, LoadingOverlay, Stack } from "@mantine/core";
import { Plus } from "../icons/Plus";
import { useState } from "react";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { getPaymentCardsByUser } from "@/api/card/getPaymentCardsByUser";
import { CreditCardModal } from "./components/CreditCardModal";
import { CreditCardItem } from "../CreditСardItem";
import { deletePaymentCard } from "@/api/card/deletePaymentCard";
import { showErrorNotification, showSuccessNotification } from "@/utils";
import { ErrorWrapper } from "@/types";
import { AxiosError } from "axios";

export const CreditCardItemList = () => {
  const [active, setActive] = useState(0);
  const [opened, setOpened] = useState(false);

  const queryClient = useQueryClient();

  const { data, isLoading } = useQuery({
    queryFn: getPaymentCardsByUser,
    queryKey: [getPaymentCardsByUser.queryKey],
  });

  const { mutate } = useMutation({
    mutationFn: deletePaymentCard,
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: [getPaymentCardsByUser.queryKey],
      });
      showSuccessNotification("Карта успешно удалена!");
    },
    onError: (error: AxiosError<{ error: ErrorWrapper }, any>) => {
      showErrorNotification(error);
    },
  });

  return (
    <Stack gap={12}>
      <CreditCardModal opened={opened} onClose={() => setOpened(false)} />

      <Button
        onClick={() => setOpened(true)}
        mih={48}
        variant="dashed"
        leftSection={<Plus fill="var(--mantine-color-accent-0)" />}
      >
        Добавить
      </Button>

      <LoadingOverlay
        visible={isLoading}
        zIndex={1}
        overlayProps={{ radius: "sm", blur: 2 }}
        loaderProps={{ color: "primary.0", type: "bars" }}
      />

      {(data?.data ?? []).map((card, ind) => (
        <CreditCardItem
          onDelete={() => mutate(card.uid)}
          onSelect={() => setActive(ind)}
          active={ind === active}
          key={ind}
        >
          {card.number}
        </CreditCardItem>
      ))}
    </Stack>
  );
};
