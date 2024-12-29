"use client";

import {
  Card,
  Text,
  Stack,
  useMatches,
  Image,
  Box,
  Overlay,
  LoadingOverlay,
} from "@mantine/core";
import { ProductItem } from "@/types";
import { memo, useCallback, useEffect, useMemo, useState } from "react";

import { ItemCardIcon } from "./ItemCardIcon";
import styles from "./ItemCard.module.css";
import { ItemCounter } from "./ItemCounter";
import { getFallbackImg, showErrorNotification } from "@/utils";
import { Trash } from "../icons/Trash";
import { Refresh } from "../icons/Refresh";
import { useMutation } from "@tanstack/react-query";
import { deleteProduct } from "@/api/products/deleteProduct";
import { AxiosError } from "axios";
import { reviveProduct } from "@/api/products/reviveProduct";

type Props = {
  item: ProductItem;
  onExtended?: () => void;
  editable?: boolean;
};

const mapTypeToValues: Record<string, any> = {
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

export const ItemCard = memo(
  ({ item, onExtended, editable = false }: Props) => {
    const type = useMatches({
      base: "small",
      md: "default",
    });
    const [deleted, setDeleted] = useState(false);

    const { maw, priceFz, priceLh, infoFz, infoLh, nameFz, nameLh } =
      mapTypeToValues[type];

    const { price, name, imgs = [], weight, ccal } = item;

    const fallbackSrc = useMemo(() => getFallbackImg(name), [name]);

    const { mutate: mutateDelete, isPending: isPendingDelete } = useMutation({
      mutationFn: deleteProduct,
      mutationKey: [deleteProduct.queryKey],
      onSuccess: () => {
        setDeleted(true);
      },
      onError: (error: AxiosError<any>) => {
        showErrorNotification(error);
      },
    });

    const { mutate: mutateRevive, isPending: isPendingRevive } = useMutation({
      mutationFn: reviveProduct,
      mutationKey: [reviveProduct.queryKey],
      onSuccess: () => {
        setDeleted(false);
      },
      onError: (error: AxiosError<any>) => {
        showErrorNotification(error);
      },
    });

    useEffect(() => {
      setDeleted(item.isDeleted);
    }, [item]);

    const handleDelete = useCallback(() => {
      mutateDelete(item.id);
    }, [item.id, deleted]);

    const handleRevive = useCallback(() => {
      mutateRevive(item.id);
    }, [item.id, deleted]);

    const isPending = isPendingDelete || isPendingRevive;

    return (
      <Card p={8} w={maw} radius="md" withBorder pos="relative">
        <Card.Section>
          <LoadingOverlay
            visible={isPending}
            overlayProps={{ radius: "sm", blur: 2 }}
            loaderProps={{ color: "accent.0", type: "bars" }}
          />

          {deleted && (
            <Overlay
              onClick={handleRevive}
              color="var(--mantine-color-accent-0)"
              backgroundOpacity={0.85}
              className={styles.overlay}
              zIndex={1}
            >
              <Refresh size={30} fill="var(--mantine-color-bg-2)" />
            </Overlay>
          )}

          {editable && (
            <Box className={styles.deleteIcon} onClick={handleDelete}>
              <Trash />
            </Box>
          )}

          <ItemCardIcon type="max" onClick={() => onExtended?.()} />

          <Image
            style={{ userSelect: "none" }}
            loading="lazy"
            src={imgs[0]}
            className={styles.img}
            alt={name}
            fit="contain"
            fallbackSrc={fallbackSrc}
            w="100%"
          />
        </Card.Section>

        <Stack mt={8} gap={4}>
          <Text lh={`${priceLh}px`} fw={700} fz={priceFz} c="accent.0">
            {price} Руб.
          </Text>
          <Text
            truncate="end"
            lh={`${infoLh}px`}
            fw={500}
            fz={infoFz}
            c="accent.2"
          >
            {weight} грамм/{ccal} ккал.
          </Text>
        </Stack>

        <Text
          truncate="end"
          lh={`${nameLh}px`}
          fw={500}
          fz={nameFz}
          mt={8}
          c="accent.0"
        >
          {name}
        </Text>

        {!editable && <ItemCounter item={item} />}
      </Card>
    );
  }
);
