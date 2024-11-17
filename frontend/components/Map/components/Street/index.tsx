import { Select } from "@mantine/core";
import { useState } from "react";
import { useMapFormContext } from "../../context";

export const Street = () => {
  const [searchValue, setSearchValue] = useState("");
  const form = useMapFormContext();

  form.watch("address", ({ value }) => {
    setSearchValue(value);
  });

  return (
    <Select
      w="100%"
      size="md"
      label="Улица"
      placeholder="Введите улицу"
      searchValue={searchValue}
      onSearchChange={setSearchValue}
      searchable
      key={form.key("address")}
      {...form.getInputProps("address")}
    />
  );
};
