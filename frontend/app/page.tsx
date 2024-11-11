"use client";

import { RecipeList } from "@/components/RecipeList";
import { YouAlreadyOrdered } from "@/components/YouAlreadOrdered";
import { YouMayLike } from "@/components/YouMayLike";
import { Stack } from "@mantine/core";

export default function HomePage() {
  return (
    <Stack gap={24} p={8} m={0} miw="100%">
      <YouAlreadyOrdered />

      <RecipeList />

      <YouMayLike />
    </Stack>
  );
}
