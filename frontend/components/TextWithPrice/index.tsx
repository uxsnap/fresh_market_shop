import { Group, Text } from "@mantine/core";
import { memo } from "react";

type TextType = "sm" | "md" | "lg";

type Props = {
  text?: string;
  price?: number;
  type?: TextType;
};

const mapTypeToText: Record<TextType, { fz: number; fw?: number }> = {
  sm: {
    fz: 14,
  },
  md: {
    fz: 18,
  },
  lg: {
    fz: 22,
    fw: 700,
  },
};

export const TextWithPrice = memo(
  ({ type = "md", text = "Какой-то текст", price = 200 }: Props) => {
    const { fz, fw = 400 } = mapTypeToText[type];

    return (
      <Group justify="space-between">
        <Text c="accent.0" fz={fz} fw={fw}>
          {text}
        </Text>

        <Text c="accent.0" fz={fz} fw="700">
          {price}₽
        </Text>
      </Group>
    );
  }
);
