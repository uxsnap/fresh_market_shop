"use client";

import { theme } from "@/theme";
import { Button, Group, useMantineColorScheme } from "@mantine/core";

export function ColorSchemeToggle() {
  const { setColorScheme } = useMantineColorScheme();

  return (
    <Group justify="center" mt="xl">
      <Button variant="accent" onClick={() => setColorScheme("light")}>
        Сохранить
      </Button>
      <Button variant="accent-reverse" onClick={() => setColorScheme("light")}>
        Тупо
      </Button>
      <Button variant="secondary" onClick={() => setColorScheme("light")}>
        Тупо
      </Button>
      <Button variant="outline" onClick={() => setColorScheme("light")}>
        Тупо
      </Button>
      <Button variant="danger" onClick={() => setColorScheme("light")}>
        Тупо
      </Button>
      <Button variant="dashed" onClick={() => setColorScheme("light")}>
        Тупо
      </Button>
    </Group>
  );
}
