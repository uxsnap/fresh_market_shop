import { Select } from "@mantine/core";
import { UseFormReturnType } from "@mantine/form";
import { useState } from "react";
import { MapForm } from "../../types";

type Props = {
  className: string;
  form: UseFormReturnType<MapForm>;
};

export const Search = ({ className, form }: Props) => {
  const [searchValue, setSearchValue] = useState("");

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
      {...form.getInputProps("address")}
    />
  );
};
