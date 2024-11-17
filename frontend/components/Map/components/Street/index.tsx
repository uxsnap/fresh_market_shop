import { useState } from "react";
import { useMapFormContext } from "../../context";
import { Select } from "@/components/Select";
import { useQuery } from "@tanstack/react-query";
import { getAddresses } from "@/api/address/getAddresses";
import { useDebouncedValue } from "@mantine/hooks";

export const Street = () => {
  const [searchValue, setSearchValue] = useState("");
  const [debounced] = useDebouncedValue(searchValue, 100);
  const form = useMapFormContext();

  // const { data } = useQuery({
  //   queryFn: () => getAddresses(),
  //   queryKey: [getAddresses.queryKey, debounced],
  //   select(data) {
  //     return data.data.map((city) => ({
  //       label: city.name,
  //       value: city.uid,
  //     }));
  //   },
  // });

  return (
    <Select
      w="100%"
      size="md"
      label="Улица"
      placeholder="Введите улицу"
      searchValue={searchValue}
      onSearchChange={setSearchValue}
      searchable
      data={[]}
      key={form.key("street")}
      {...form.getInputProps("street")}
    />
  );
};
