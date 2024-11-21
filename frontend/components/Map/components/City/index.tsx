import { useState } from "react";
import { useMapFormContext } from "../../context";
import { useQuery } from "@tanstack/react-query";
import { getCities } from "@/api/address/getCities";
import { Select } from "@mantine/core";

export const City = () => {
  const [searchValue, setSearchValue] = useState("");
  const form = useMapFormContext();

  const { data } = useQuery({
    queryFn: getCities,
    queryKey: [getCities.queryKey],
    select(data) {
      return data.data.map((city) => ({
        label: city.name,
        value: city.uid,
      }));
    },
    staleTime: Infinity,
  });

  return (
    <Select
      w="100%"
      size="md"
      label="Город"
      placeholder="Выберите город"
      searchValue={searchValue}
      onSearchChange={setSearchValue}
      searchable
      key={form.key("city")}
      data={data ?? []}
      allowDeselect={false}
      withAsterisk
      {...form.getInputProps("city")}
    />
  );
};
