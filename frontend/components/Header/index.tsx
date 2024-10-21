import {
  Button,
  Container,
  Flex,
  Group,
  Stack,
  TextInput,
} from "@mantine/core";
import { Menu } from "../icons/Menu";
import { Glass } from "../icons/Glass";
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
  <Stack mx="auto" className={styles.root} maw={1454}>
    <Flex align="center" justify="space-between">
      <Group w="100%" wrap="nowrap">
        <Button onClick={onNavbar} h={38} w={38} px={8} variant="secondary">
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
