import { Group, NumberInput, Stack } from "@mantine/core";

import { memo } from "react";
import { useMapFormContext } from "../../context";

export const AdditionalFieldsForm = memo(() => {
  const form = useMapFormContext();

  return (
    <Group gap={20} w="100%" wrap="nowrap">
      <Stack gap={12} w="100%">
        <NumberInput
          w="100%"
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
          w="100%"
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
      </Stack>

      <Stack gap={12} w="100%">
        <NumberInput
          w="100%"
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
          w="100%"
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
      </Stack>
    </Group>
  );
});
