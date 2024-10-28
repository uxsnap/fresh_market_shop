import { Button, Group, Stack } from "@mantine/core";
import { Menu } from "../icons/Menu";
import { Location } from "../Location";
import { DeliveryTime } from "../DeliveryTime";
import { User } from "../icons/User";
import styles from "./Header.module.css";
import { CartButton } from "./components/CartButton";
import { Search } from "./components/Search";
import { memo } from "react";
import { Auth } from "../Auth";
import { useDisclosure } from "@mantine/hooks";

type Props = {
  onNavbar: () => void;
};

export const Header = memo(({ onNavbar }: Props) => {
  const [opened, { open, close }] = useDisclosure(false);

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
          <Button
            onClick={onNavbar}
            h={38}
            miw={38}
            maw={38}
            px={8}
            variant="secondary"
          >
            <Menu size={24} />
          </Button>

          <Search className={styles.topSearch} />

          <Location />

          <DeliveryTime />
        </Group>

        <Group wrap="nowrap" gap={24} align="center">
          <CartButton />

          <Button
            className={styles.userButton}
            h={24}
            w={24}
            p={0}
            variant="outline"
            onClick={open}
          >
            <User />
          </Button>

          <Auth close={close} opened={opened} />
        </Group>
      </Group>

      <Search maw="100%" className={styles.bottomSearch} />
    </Stack>
  );
});
