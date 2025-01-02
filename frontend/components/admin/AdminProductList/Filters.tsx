import { Group, TextInput } from "@mantine/core";
import { useDebouncedCallback } from "@mantine/hooks";
import { ChangeEvent, useState } from "react";

type Filters = {
  name: string;
};

type Props = {
  setFilters: (filters: Filters) => void;
};

export const Filters = ({ setFilters }: Props) => {
  const [innerFilters, setInnerFilters] = useState({
    name: "",
  });

  const handleFilters = useDebouncedCallback((filters: Filters) => {
    console.log("here");

    setFilters(filters);
  }, 200);

  const handleName = (event: ChangeEvent) => {
    const newFilters = {
      ...innerFilters,
      name: (event.target as HTMLInputElement).value,
    };

    setInnerFilters(newFilters);
    handleFilters(newFilters);
  };

  return (
    <Group pl={16} gap={12}>
      <TextInput
        miw={200}
        label="Имя"
        placeholder="Введите имя"
        onChange={handleName}
        value={innerFilters.name}
      />
    </Group>
  );
};
