import { Select } from "@mantine/core";
import { useState } from "react";
import { useMapFormContext } from "../../context";

type Props = {
  className: string;
};

export const Search = ({ className }: Props) => {
  const [searchValue, setSearchValue] = useState("");
  const form = useMapFormContext();

  form.watch("address", ({ value }) => {
    setSearchValue(value);
  });

  return (
    <Select
      className={className}
      size="md"
      label="Адрес"
      placeholder="Выберите адрес"
      searchValue={searchValue}
      onSearchChange={setSearchValue}
      searchable
      limit={5}
      key={form.key("address")}
      {...form.getInputProps("address")}
    />
  );
};
