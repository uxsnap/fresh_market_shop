"use client";

import { theme } from "@/theme";
import { Button, Group, useMantineColorScheme } from "@mantine/core";

export function ColorSchemeToggle() {
  const { setColorScheme } = useMantineColorScheme();

  return (
    <Group justify="center" mt="xl">
      <Button color="secondary.0" onClick={() => setColorScheme("light")}>
        Сохранить
      </Button>
      {/* <Button onClick={() => setColorScheme("dark")}>Dark</Button>
      <Button onClick={() => setColorScheme("auto")}>Auto</Button> */}
    </Group>
  );
}
