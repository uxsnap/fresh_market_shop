import { Modal, Title } from "@mantine/core";
import { memo } from "react";

import styles from "./ModalHeader.module.css";

export const ModalHeader = () => (
  <Modal.Header className={styles.header}>
    <Title c="accent.0">Укажите ваш адрес</Title>
    <Modal.CloseButton size="32px" c="accent.0" />
  </Modal.Header>
);
