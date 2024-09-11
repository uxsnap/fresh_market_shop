"use client";

import { RecipeModal } from "@/components/RecipeModal";
import { Container } from "@mantine/core";

export default function HomePage() {
  return (
    <Container p={10}>
      <RecipeModal />
    </Container>
  );
}
