import { logoutUser } from "@/api/auth/logout";
import { Avatar } from "@/components/Avatar";
import { useAuthStore } from "@/store/auth";
import cn from "classnames";
import { Box, Divider, LoadingOverlay, Popover, Stack } from "@mantine/core";
import { useMutation } from "@tanstack/react-query";
import { MouseEvent } from "react";

import styles from "./UserMenu.module.css";
import Link from "next/link";

export const UserMenu = () => {
  const setLogged = useAuthStore((s) => s.setLogged);

  const { mutate, isPending } = useMutation({
    mutationFn: logoutUser,
    onSuccess: () => {
      setLogged(false);
    },
  });

  const handleLogout = (e: MouseEvent<HTMLElement>) => {
    mutate();
  };

  return (
    <Box mah={38} maw={38} pos="relative">
      <Popover
        clickOutsideEvents={["mouseup", "touchend"]}
        position="bottom"
        radius="md"
      >
        <Popover.Target>
          <Box style={{ cursor: "pointer" }}>
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
            <Link className={styles.link} href="/profile">
              Перейти в профиль
            </Link>

            <Divider mx={12} size="xs" bg="var(--mantine-color-accent-0)" />

            <Link
              href="#"
              className={cn(styles.link, styles.bottom)}
              onClick={handleLogout}
            >
              Выйти из системы
            </Link>
          </Stack>
        </Popover.Dropdown>
      </Popover>
    </Box>
  );
};
