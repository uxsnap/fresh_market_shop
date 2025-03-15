import { logoutUser } from "@/api/auth/logout";
import { Avatar } from "@/components/Avatar";
import { useAuthStore } from "@/store/auth";
import cn from "classnames";
import { Box, Divider, LoadingOverlay, Popover, Stack } from "@mantine/core";
import { useMutation } from "@tanstack/react-query";
import { useState } from "react";

import styles from "./UserMenu.module.css";
import Link from "next/link";
import { useRouter } from "next/navigation";

export const UserMenu = () => {
  const setLogged = useAuthStore((s) => s.setLogged);
  const admin = useAuthStore((s) => s.admin);
  const [opened, setOpened] = useState(false);
  const router = useRouter();

  const { mutate, isPending } = useMutation({
    mutationFn: logoutUser,
    onSuccess: () => {
      setLogged(false);
      router.push("/");
    },
  });

  const handleLogout = () => {
    mutate();
    setOpened(false);
  };

  return (
    <Box mah={38} maw={38} pos="relative">
      <Popover
        opened={opened}
        closeOnClickOutside
        onChange={setOpened}
        onClose={() => setOpened(false)}
        position="bottom"
        radius="md"
      >
        <Popover.Target>
          <Box onClick={() => setOpened(!opened)} style={{ cursor: "pointer" }}>
            <Avatar size="small" />
          </Box>
        </Popover.Target>

        <Popover.Dropdown p={0}>
          <LoadingOverlay
            visible={isPending}
            zIndex={1}
            overlayProps={{ radius: "xs", blur: 3 }}
            loaderProps={{ color: "primary.0", type: "bars" }}
          />

          <Stack gap={0}>
            <Link
              prefetch
              className={styles.link}
              onClick={() => setOpened(false)}
              href="/profile"
            >
              <div>Перейти в профиль</div>
            </Link>

            <Divider mx={12} size="xs" bg="var(--mantine-color-accent-0)" />

            {admin && (
              <Link className={styles.link} href="/admin">
                <div>Перейти в панель админа</div>
              </Link>
            )}

            <Divider mx={12} size="xs" bg="var(--mantine-color-accent-0)" />

            <Link
              href="#"
              className={cn(styles.link, styles.bottom)}
              onClick={handleLogout}
            >
              <div>Выйти из системы</div>
            </Link>
          </Stack>
        </Popover.Dropdown>
      </Popover>
    </Box>
  );
};
