import { Button, Stack } from "@mantine/core";
import { AddressItem } from "../AddressItem";
import { Plus } from "../icons/Plus";
import { useState } from "react";

export const AddressItemList = () => {
  const [active, setActive] = useState(0);

  return (
    <Stack gap={12}>
      <Button
        mih={48}
        variant="dashed"
        leftSection={<Plus fill="var(--mantine-color-accent-0)" />}
      >
        Добавить
      </Button>

      {Array.from({ length: 5 }).map((_, ind) => (
        <AddressItem
          onSelect={() => setActive(ind)}
          active={ind === active}
          key={ind}
        />
      ))}
    </Stack>
  );
};
