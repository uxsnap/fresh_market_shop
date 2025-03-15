import { Box, LoadingOverlay, Modal } from "@mantine/core";

import styles from "./Map.module.css";

import { MapFields } from "./components/MapFields";
import { YmapsWrapper } from "./components/YmapsWrapper";
import { MapForm, MapFormProvider, useMapForm } from "./context";
import { BottomCards } from "./components/BottomCards";
import { ModalHeader } from "./components/ModalHeader";
import { isNotEmpty } from "@mantine/form";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { addDeliveryAddress } from "@/api/user/addDeliveryAddress";
import { showErrorNotification, showSuccessNotification } from "@/utils";
import { AxiosError } from "axios";
import { ErrorWrapper } from "@/types";
import { getDeliveryAddresses } from "@/api/user/getDeliveryAddresses";
import { useMapStore } from "@/store/map";

export const Map = () => {
  const queryClient = useQueryClient();

  const isMapOpen = useMapStore((s) => s.isMapOpen);
  const setIsMapOpen = useMapStore((s) => s.setIsMapOpen);
  const setMapAddress = useMapStore((s) => s.setMapAddress);

  const handleClose = () => {
    close();
    setMapAddress();
    setIsMapOpen(false);
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
    mutationFn: addDeliveryAddress,
    mutationKey: [addDeliveryAddress.queryKey],
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: [getDeliveryAddresses.queryKey],
      });
      showSuccessNotification("Адрес успешно добавлен!");
      setIsMapOpen(false);
    },
    onError: (error: AxiosError<{ error: ErrorWrapper }, any>) => {
      showErrorNotification(error);
    },
  });

  const handleSubmit = form.onSubmit((values) => {
    const submitValues = Object.keys(values).reduce((acc, cur) => {
      const key = cur as keyof MapForm;

      if (!values[key]) {
        return acc;
      }

      acc[cur] = values[key];

      return acc;
    }, {} as any);

    mutate(submitValues);
  });

  return (
    <Modal.Root
      classNames={{
        inner: styles.inner,
      }}
      radius={0}
      fullScreen
      opened={isMapOpen}
      onClose={handleClose}
    >
      <Modal.Content className={styles.content}>
        <ModalHeader />

        <Modal.Body p={0} h="100%">
          <MapFormProvider form={form}>
            <form className={styles.form} onSubmit={handleSubmit}>
              <YmapsWrapper />

              <Box className={styles.mainBox}>
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
