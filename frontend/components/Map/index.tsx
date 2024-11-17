import { Modal } from "@mantine/core";

import styles from "./Map.module.css";

import { MapFields } from "./components/MapFields";
import { YmapsWrapper } from "./components/YmapsWrapper";
import { MapFormProvider, useMapForm } from "./context";
import { BottomCards } from "./components/BottomCards";
import { ModalHeader } from "./components/ModalHeader";

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
      city: "",
      street: "",
    },
  });

  const handleSubmit = form.onSubmit((values) => {
    console.log(values);
  });

  return (
    <Modal.Root
      classNames={{
        inner: styles.inner,
      }}
      radius={0}
      fullScreen
      opened={opened}
      onClose={handleClose}
    >
      <Modal.Content className={styles.content}>
        <ModalHeader />

        <Modal.Body p={0} h="100%">
          <MapFormProvider form={form}>
            <form className={styles.form} onSubmit={handleSubmit}>
              <YmapsWrapper />

              <MapFields />

              <BottomCards items={[]} />
            </form>
          </MapFormProvider>
        </Modal.Body>
      </Modal.Content>
    </Modal.Root>
  );
};
