import { RangeSlider } from "@mantine/core";

export const Slider = () => {
  return (
    <RangeSlider
      color="primary.1"
      min={0}
      max={1000}
      mt={40}
      defaultValue={[500, 600]}
      radius="sm"
      size="xs"
      marks={[
        { value: 100, label: "100" },
        { value: 500, label: "500" },
        { value: 800, label: "700" },
      ]}
      labelAlwaysOn
    />
  );
};
