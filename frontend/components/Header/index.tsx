import { Burger, Group, Stack, Image } from "@mantine/core";
import { Location } from "../Location";
import styles from "./Header.module.css";
import { CartButton } from "./components/CartButton";
import { Search } from "./components/Search";
import { memo } from "react";
import { UserAuth } from "./components/UserAuth";
import { useDisclosure } from "@mantine/hooks";
import { useRouter } from "next/navigation";

type Props = {
  onNavbar: () => void;
  opened: boolean;
};

export const Header = memo(({ opened, onNavbar }: Props) => {
  const router = useRouter();

  const handleNavbar = () => {
    onNavbar();
  };

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

          <Burger
            h={38}
            w={38}
            className={styles.burger}
            opened={opened}
            onClick={handleNavbar}
            aria-label="Toggle navigation"
          />

          <Search className={styles.topSearch} />

          <Location />

          {/* <DeliveryTime /> */}
        </Group>

        <UserAuth>
          <CartButton />
        </UserAuth>
      </Group>

      <Search maw="100%" className={styles.bottomSearch} />
    </Stack>
  );
});
