"use client";

import { Group, TextInput } from "@mantine/core";
import { Glass } from "../icons/Glass";

export function ColorSchemeToggle() {
  return (
    <Group justify="center" mt="xl">
      <TextInput
        leftSection={<Glass size={16} />}
        placeholder="Yesss"
        error="Dadad"
        label="Test"
      />
    </Group>
  );
}
