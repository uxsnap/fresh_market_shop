import { Group, NumberInput, Stack } from "@mantine/core";

import { memo } from "react";
import { useMapStore } from "@/store/map";
import { useMapFormContext } from "../../context";

export const AdditionalFieldsForm = memo(() => {
  const form = useMapFormContext();

  const fields = useMapStore((s) => s.fields);
  const setFields = useMapStore((s) => s.setFields);

  form.watch("addressUid", () => {
    setFields({
      apartment: undefined,
      entrance: undefined,
      floor: undefined,
      code: undefined,
    });
  });

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
          onChange={(v) =>
            setFields({ ...fields, apartment: parseInt(v.toString(), 10) })
          }
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
          onChange={(v) =>
            setFields({ ...fields, entrance: parseInt(v.toString(), 10) })
          }
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
          onChange={(v) =>
            setFields({ ...fields, floor: parseInt(v.toString(), 10) })
          }
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
          onChange={(v) =>
            setFields({ ...fields, code: parseInt(v.toString(), 10) })
          }
        />
      </Stack>
    </Group>
  );
});
