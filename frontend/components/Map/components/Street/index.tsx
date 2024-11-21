import { useState } from "react";
import { useMapFormContext } from "../../context";
import { useQuery } from "@tanstack/react-query";
import { getAddresses } from "@/api/address/getAddresses";
import { useDebouncedValue } from "@mantine/hooks";
import { Select } from "@mantine/core";
import { useMapStore } from "@/store/map";

export const Street = () => {
  const [curCity, setCurCity] = useState("");

  const searchValue = useMapStore((s) => s.searchValue);
  const setSearchValue = useMapStore((s) => s.setSearchValue);

  const [debounced] = useDebouncedValue(searchValue, 200);
  const form = useMapFormContext();

  form.watch("city", ({ value }) => {
    setCurCity(value);
  });

  const { data } = useQuery({
    queryFn: () => getAddresses(curCity, debounced),
    queryKey: [getAddresses.queryKey, debounced],
    enabled: !!debounced.length,
    select(data) {
      return data.data.map((s) => ({
        label: `${s.street} ${s.houseNumber}`,
        value: s.uid,
      }));
    },
  });

  return (
    <Select
      disabled={!curCity}
      w="100%"
      size="md"
      label="Улица"
      placeholder="Введите улицу"
      searchValue={searchValue}
      onSearchChange={setSearchValue}
      searchable
      data={data ?? []}
      nothingFoundMessage="Ничего не найдено"
      allowDeselect={false}
      filter={({ options }) => options}
      key={form.key("addressUid")}
      withAsterisk
      {...form.getInputProps("addressUid")}
    />
  );
};