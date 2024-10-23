"use client";

import { YouMayLike } from "@/components/YouMayLike";
import { Container } from "@mantine/core";

export default function HomePage() {
  return (
    <Container p={8} m={0} miw="100%">
      <YouMayLike />
    </Container>
  );
}
