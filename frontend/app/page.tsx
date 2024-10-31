"use client";

import { RecipeList } from "@/components/RecipeList";
import { YouMayLike } from "@/components/YouMayLike";
import { Stack } from "@mantine/core";

export default function HomePage() {
  return (
    <Stack gap={24} p={8} m={0} miw="100%">
      <RecipeList />

      <YouMayLike />
    </Stack>
  );
}
