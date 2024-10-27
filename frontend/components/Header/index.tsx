import { Button, Flex, Group, Stack } from "@mantine/core";
import { Menu } from "../icons/Menu";
import { Location } from "../Location";
import { DeliveryTime } from "../DeliveryTime";
import { User } from "../icons/User";
import styles from "./Header.module.css";
import { CartButton } from "./components/CartButton";
import { Search } from "./components/Search";

type Props = {
  onNavbar: () => void;
};

export const Header = ({ onNavbar }: Props) => (
  <Stack justify="center" className={styles.root}>
    <Flex mx="auto" align="center" justify="space-between" w="100%" maw={1520}>
      <Group w="100%" wrap="nowrap">
        <Button onClick={onNavbar} h={38} miw={38} maw={38} px={8} variant="secondary">
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
        >
          <User />
        </Button>
      </Group>
    </Flex>

    <Search maw="100%" className={styles.bottomSearch} />
  </Stack>
);
