import styles from "./MapFields.module.css";
import { Button, Group, Stack } from "@mantine/core";
import { Search } from "../Search";
import { AdditionalFieldsForm } from "../AdditionalFieldsForm";
import { memo } from "react";

export const MapFields = memo(() => {
  return (
    <Stack className={styles.fieldsWrapper}>
      <Group
        className={styles.addressWrapper}
        grow
        gap={16}
        w="100%"
        align="end"
        wrap="nowrap"
      >
        <Search className={styles.address} />

        <Button
          type="submit"
          className={styles.button}
          h={42}
          px={4}
          fz={14}
          maw={150}
          variant="accent"
        >
          Добавить адрес
        </Button>
      </Group>

      <AdditionalFieldsForm />
    </Stack>
  );
});
