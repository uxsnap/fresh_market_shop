import { Group, Stack, Image, Title } from "@mantine/core";
import styles from "./Header.module.css";
import { memo } from "react";
import { useRouter } from "next/navigation";
import { UserAuth } from "@/components/Header/components/UserAuth";

export const AdminHeader = memo(() => {
  const router = useRouter();

  return (
    <Stack justify="center" className={styles.root}>
      <Group
        gap={16}
        mx="auto"
        align="center"
        justify="space-between"
        w="100%"
        wrap="nowrap"
        maw={1520}
      >
        <Group w="100%" wrap="nowrap">
          <Image
            className={styles.logo}
            onClick={() => router.push("/")}
            src="Logo.svg"
            alt="_Logo"
          />

          <Title order={2} c="accent.0">
            ADMIN PANEL
          </Title>
        </Group>

        <UserAuth />
      </Group>
    </Stack>
  );
});
