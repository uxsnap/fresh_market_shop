import { Modal, Title } from "@mantine/core";

import styles from "./Map.module.css";

import { MapFields } from "./components/MapFields";
import { YmapsWrapper } from "./components/YmapsWrapper";
import { MapFormProvider, useMapForm } from "./context";

type Props = {
  opened?: boolean;
  onClose: () => void;
};

export const Map = ({ opened = false, onClose }: Props) => {
  const handleClose = () => {
    close();
    onClose();
  };

  const form = useMapForm({
    mode: "uncontrolled",
    initialValues: {
      address: "",
    },
  });

  const handleSubmit = form.onSubmit((values) => {
    console.log(values);
  });

  return (
    <Modal.Root opened={opened} onClose={handleClose}>
      <Modal.Overlay />

      <Modal.Content className={styles.content}>
        <Modal.Header className={styles.header}>
          <Title c="accent.0"> Укажите ваш адрес</Title>
          <Modal.CloseButton size="32px" c="accent.0" />
        </Modal.Header>

        <Modal.Body px={20} py={12}>
          <MapFormProvider form={form}>
            <form onSubmit={handleSubmit}>
              <MapFields />

              <YmapsWrapper />
            </form>
          </MapFormProvider>
        </Modal.Body>
      </Modal.Content>
    </Modal.Root>
  );
};
