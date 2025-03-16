"use client";

import styles from "./MapFields.module.css";
import { Button, Modal, Stack, useMatches } from "@mantine/core";
import { City } from "../City";
import { AdditionalFieldsForm } from "../AdditionalFieldsForm";
import { memo, PropsWithChildren, useEffect } from "react";
import { Street } from "../Street";
import { useMapStore } from "@/store/map";
import { ModalHeader } from "../ModalHeader";

export const MapFields = () => {
  const isFieldsModalOpen = useMapStore((s) => s.isFieldsModalOpen);
  const setIsFieldsModalOpen = useMapStore((s) => s.setIsFieldsModalOpen);

  const modalWrapper = useMatches({
    base: true,
    md: false,
  });

  useEffect(() => {
    setIsFieldsModalOpen(false);
  }, [modalWrapper]);

  const Wrapper = ({ children }: PropsWithChildren) => {
    if (modalWrapper) {
      return (
        <Modal.Root
          fullScreen
          opened={isFieldsModalOpen}
          onClose={() => setIsFieldsModalOpen(false)}
        >
          <Modal.Content className={styles.content}>
            <ModalHeader />

            <Modal.Body className={styles.body}>{children}</Modal.Body>
          </Modal.Content>
        </Modal.Root>
      );
    }

    return <>{children}</>;
  };

  return (
    <Wrapper>
      <Stack justify="space-between" className={styles.root}>
        <Stack className={styles.addressWrapper} gap={16} w="100%" align="end">
          <City />

          <Street />

          <AdditionalFieldsForm />
        </Stack>

        <Button type="submit" h={48} fz={18} variant="accent">
          Сохранить адрес
        </Button>
      </Stack>
    </Wrapper>
  );
};
