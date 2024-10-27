import Link from "next/link";
import { ArrowLeft } from "../icons/ArrowLeft";
import { Box, Group, Title } from "@mantine/core";

import styles from "./BackToCatalog.module.css";
import { RemoveAll } from "../pages/cart/RemoveAll";

type Props = {
  empty?: boolean;
};

export const BackToCatalog = ({ empty = true }: Props) => (
  <Group justify="space-between">
    <Link className={styles.root} href={"/"}>
      <Group gap={8}>
        <ArrowLeft fill="var(--mantine-color-accent-0)" />
        <Title className={styles.title} c="accent.0">
          Вернуться обратно к каталогу
        </Title>
      </Group>
    </Link>

    {!empty && (
      <Box hiddenFrom="md">
        <RemoveAll />
      </Box>
    )}
  </Group>
);
