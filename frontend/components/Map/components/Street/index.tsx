import { useEffect, useMemo } from "react";
import { useMapFormContext } from "../../context";
import { useQuery } from "@tanstack/react-query";
import { getAddresses } from "@/api/address/getAddresses";
import { useDebouncedValue } from "@mantine/hooks";
import { Select } from "@mantine/core";
import { useMapStore } from "@/store/map";
import { getStreetAndHouseNumber } from "@/utils";

export const Street = () => {
  const searchValue = useMapStore((s) => s.searchValue);
  const setSearchValue = useMapStore((s) => s.setSearchValue);
  const curCity = useMapStore((s) => s.city);

  const setMapAddress = useMapStore((s) => s.setMapAddress);

  const [debounced] = useDebouncedValue(searchValue, 200);
  const form = useMapFormContext();

  const fetchAddresses = () => {
    const [name, houseNumber] = getStreetAndHouseNumber(debounced);
    return getAddresses(curCity, name, houseNumber);
  };

  const { data } = useQuery({
    queryFn: fetchAddresses,
    queryKey: [getAddresses.queryKey, curCity, debounced],
    enabled: !!debounced.length,
  });

  form.watch("addressUid", ({ value }) => {
    const curMapActiveAddress = data?.data.find((a) => a.uid === value);

    setMapAddress(curMapActiveAddress);
  });

  useEffect(() => {
    if (!curCity || !data?.data.length) {
      return;
    }

    form.setFieldValue("addressUid", data.data[0].uid);
  }, [curCity, data?.data]);

  const preparedData = useMemo(() => {
    return data?.data.map((a) => ({
      label: `${a.street} ${a.houseNumber}`,
      value: a.uid,
    }));
  }, [data]);

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
      data={preparedData ?? []}
      nothingFoundMessage="Ничего не найдено"
      allowDeselect={false}
      filter={({ options }) => options}
      key={form.key("addressUid")}
      withAsterisk
      {...form.getInputProps("addressUid")}
    />
  );
};
