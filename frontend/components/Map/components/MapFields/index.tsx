import styles from "./MapFields.module.css";
import { Button, Stack } from "@mantine/core";
import { City } from "../City";
import { AdditionalFieldsForm } from "../AdditionalFieldsForm";
import { memo } from "react";
import { Street } from "../Street";

export const MapFields = memo(() => (
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
));
