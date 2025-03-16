import { useState } from "react";
import { useMapFormContext } from "../../context";
import { useQuery } from "@tanstack/react-query";
import { getCities } from "@/api/address/getCities";
import { Select } from "@mantine/core";
import { useMapStore } from "@/store/map";

export const City = () => {
  const [searchValue, setSearchValue] = useState("");
  const form = useMapFormContext();
  const setCurCity = useMapStore((s) => s.setCity);

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

  const props = form.getInputProps("city");

  const handleChange = (v: string) => {
    props.onChange(v);
    setCurCity(v);
  };

  return (
    <Select
      w="100%"
      size="md"
      label="Город"
      placeholder="Выберите город"
      searchValue={searchValue}
      onSearchChange={setSearchValue}
      searchable
      data={data ?? []}
      allowDeselect={false}
      withAsterisk
      key={form.key("city")}
      {...props}
      value={props.defaultValue}
      onChange={(v) => handleChange(v + "")}
    />
  );
};
