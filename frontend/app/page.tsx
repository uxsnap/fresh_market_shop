"use client";

import { YouMayLike } from "@/components/pages/home/YouMayLike";
import { Container } from "@mantine/core";

export default function HomePage() {
  return (
    <Container p={10}>
      <YouMayLike />
    </Container>
  );
}
