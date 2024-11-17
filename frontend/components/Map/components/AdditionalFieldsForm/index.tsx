import { Flex, NumberInput } from "@mantine/core";

import styles from "./AdditionalFieldsForm.module.css";
import { memo } from "react";
import { useMapFormContext } from "../../context";

export const AdditionalFieldsForm = memo(() => {
  const form = useMapFormContext();

  return (
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
        key={form.key("flat")}
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
        key={form.key("entrance")}
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
        key={form.key("floor")}
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
        key={form.key("code")}
        {...form.getInputProps("code")}
      />
    </Flex>
  );
});
