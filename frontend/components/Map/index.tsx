import { Box, LoadingOverlay, Modal, Stack } from "@mantine/core";

import styles from "./Map.module.css";

import { MapFields } from "./components/MapFields";
import { YmapsWrapper } from "./components/YmapsWrapper";
import { MapFormProvider, useMapForm } from "./context";
import { BottomCards } from "./components/BottomCards";
import { ModalHeader } from "./components/ModalHeader";
import { isNotEmpty } from "@mantine/form";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { addUserAddress } from "@/api/user/addUserAddress";
import { showErrorNotification, showSuccessNotification } from "@/utils";
import { AxiosError } from "axios";
import { ErrorWrapper } from "@/types";
import { getUserAddresses } from "@/api/user/getUserAdresses";

type Props = {
  opened?: boolean;
  onClose: () => void;
};

export const Map = ({ opened = false, onClose }: Props) => {
  const queryClient = useQueryClient();

  const handleClose = () => {
    close();
    onClose();
  };

  const form = useMapForm({
    mode: "uncontrolled",
    initialValues: {
      city: "",
      addressUid: "",
    },
    validate: {
      city: isNotEmpty("Необходимо выбрать город!"),
      addressUid: isNotEmpty("Необходимо выбрать улицу!"),
    },
  });

  const { mutate, isPending } = useMutation({
    mutationFn: addUserAddress,
    mutationKey: [addUserAddress.queryKey],
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: [getUserAddresses.queryKey],
      });
      showSuccessNotification("Адрес успешно добавлен!");
      onClose();
    },
    onError: (error: AxiosError<{ error: ErrorWrapper }, any>) => {
      showErrorNotification(error);
    },
  });

  const handleSubmit = form.onSubmit((values) => {
    mutate(values);
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

              <Box h="100%" pos="relative">
                <LoadingOverlay
                  visible={isPending}
                  zIndex={1}
                  overlayProps={{ radius: "sm", blur: 2 }}
                  loaderProps={{ color: "primary.0", type: "bars" }}
                />
                <MapFields />

                <BottomCards />
              </Box>
            </form>
          </MapFormProvider>
        </Modal.Body>
      </Modal.Content>
    </Modal.Root>
  );
};
