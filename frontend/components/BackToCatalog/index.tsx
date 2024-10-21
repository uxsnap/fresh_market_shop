import Link from "next/link";
import { ArrowLeft } from "../icons/ArrowLeft";
import { Group, Title } from "@mantine/core";

export const BackToCatalog = () => (
  <Link href={"/"} style={{ textDecoration: "none" }}>
    <Group gap={8}>
      <ArrowLeft fill="var(--mantine-color-accent-0)" />
      <Title c="accent.0" order={3}>
        Вернуться обратно к каталогу
      </Title>
    </Group>
  </Link>
);
