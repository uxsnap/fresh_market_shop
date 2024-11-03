import { User } from "@/components/icons/User";
import { Box, Button, Group } from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";

import styles from "./UserAuth.module.css";
import { Auth } from "@/components/Auth";
import { useAuthStore } from "@/store/auth";
import { Avatar } from "@/components/Avatar";
import { PropsWithChildren } from "react";

export const UserAuth = ({ children }: PropsWithChildren) => {
  const logged = useAuthStore((s) => s.logged);
  const setModalOpen = useAuthStore((s) => s.setModalOpen);

  return (
    <Group wrap="nowrap" gap={logged ? 12 : 24} align="center">
      {children}

      {logged ? (
        <Box mah={38} maw={38} hidden={!logged}>
          <Avatar size="small" />
        </Box>
      ) : (
        <Button
          hidden={logged}
          className={styles.userButton}
          h={24}
          w={24}
          p={0}
          variant="outline"
          onClick={() => setModalOpen("login")}
        >
          <User />
        </Button>
      )}

      <Auth />
    </Group>
  );
};
