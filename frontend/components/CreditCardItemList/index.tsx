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
import { useOrderStore } from "@/store/order";

export const CreditCardItemList = () => {
  const [opened, setOpened] = useState(false);

  const creditCard = useOrderStore((s) => s.creditCard);
  const setCreditCard = useOrderStore((s) => s.setCreditCard);

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
          onSelect={() => setCreditCard(card)}
          active={creditCard?.uid === card.uid}
          key={ind}
        >
          {card.number}
        </CreditCardItem>
      ))}
    </Stack>
  );
};
