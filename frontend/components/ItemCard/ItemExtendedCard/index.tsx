"use client";

import { Text, Stack, Group, Modal } from "@mantine/core";
import { memo, useCallback } from "react";

import styles from "./ItemExtendedCard.module.css";
import { ItemCardCarousel } from "../ItemCardCarousel";
import { ItemCounter } from "../ItemCounter";
import { ItemCardIcon } from "../ItemCardIcon";
import { useProductStore } from "@/store/product";

export const ItemCardExtended = memo(() => {
  const curItem = useProductStore((s) => s.curItem);
  const setCurItem = useProductStore((s) => s.setCurItem);

  const handleClose = () => setCurItem();

  const { price, name, imgs = [], description, ccal } = curItem || {};

  return (
    <Modal.Root
      p={0}
      centered
      opened={!!curItem}
      onClose={handleClose}
      closeOnClickOutside
    >
      <Modal.Overlay />

      <Modal.Content>
        <Modal.Header className={styles.header}>
          <Group w="100%" wrap="nowrap" align="center" justify="space-between">
            <Text lh={1} truncate="end" fw={700} fz={32} c="accent.0">
              {name}
            </Text>

            <ItemCardIcon type="min" onClick={handleClose} />
          </Group>
        </Modal.Header>

        <Modal.Body p={0}>
          <ItemCardCarousel
            className={styles.img}
            imgs={imgs.map((img) => img.path)}
            name={name ?? ""}
          />

          <Stack p={12} gap={12} mt={12}>
            <Stack gap={8}>
              <Group align="flex-end" gap={12}>
                <Text lh="26px" truncate="end" fw={700} fz={22} c="accent.0">
                  Описание товара
                </Text>

                <Text lh="23px" truncate="end" fw={700} fz={18} c="accent.2">
                  {ccal} ккал/порция
                </Text>
              </Group>

              <Text lh={1} fz={14} c="accent.0">
                {description}
              </Text>
            </Stack>

            <ItemCounter item={curItem}>
              <Text fw="bold" fz={18}>
                {price} ₽
              </Text>
            </ItemCounter>
          </Stack>
        </Modal.Body>
      </Modal.Content>
    </Modal.Root>
  );
});
