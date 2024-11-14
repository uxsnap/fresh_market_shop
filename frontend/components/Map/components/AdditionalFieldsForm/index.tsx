import { Flex, NumberInput } from "@mantine/core";

import styles from "./Form.module.css";
import { memo, PropsWithChildren } from "react";
import { MapForm } from "../../types";
import { UseFormReturnType } from "@mantine/form";

export const AdditionalFieldsForm = memo(
  ({ form }: PropsWithChildren<{ form: UseFormReturnType<MapForm> }>) => (
    <Flex className={styles.fields}>
      <NumberInput
        min={1}
        hideControls
        allowLeadingZeros={false}
        allowNegative={false}
        allowDecimal={false}
        lh={1}
        size="md"
        label="Квартира"
        placeholder="Введите квартиру"
        {...form.getInputProps("flat")}
      />

      <NumberInput
        min={1}
        hideControls
        allowLeadingZeros={false}
        allowNegative={false}
        allowDecimal={false}
        lh={1}
        size="md"
        label="Подъезд"
        placeholder="Введите подъезд"
        {...form.getInputProps("entrance")}
      />

      <NumberInput
        min={1}
        hideControls
        allowLeadingZeros={false}
        allowNegative={false}
        allowDecimal={false}
        lh={1}
        size="md"
        label="Этаж"
        placeholder="Введите этаж"
        {...form.getInputProps("floor")}
      />

      <NumberInput
        min={1}
        hideControls
        allowLeadingZeros={false}
        allowNegative={false}
        allowDecimal={false}
        lh={1}
        size="md"
        label="Домофон"
        placeholder="Введите домофон"
        {...form.getInputProps("code")}
      />
    </Flex>
  )
);
