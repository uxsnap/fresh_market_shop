import { User } from "@/components/icons/User";
import { Button, Group } from "@mantine/core";

import styles from "./UserAuth.module.css";
import { Auth } from "@/components/Auth";
import { useAuthStore } from "@/store/auth";
import { PropsWithChildren } from "react";
import { UserMenu } from "./UserMenu";

export const UserAuth = ({ children }: PropsWithChildren) => {
  const logged = useAuthStore((s) => s.logged);
  const setModalOpen = useAuthStore((s) => s.setModalOpen);

  if (logged === undefined) {
    return;
  }

  return (
    <Group wrap="nowrap" gap={logged ? 12 : 24} align="center">
      {children}

      {logged ? (
        <UserMenu />
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
