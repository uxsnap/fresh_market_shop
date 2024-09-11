import { Button, Flex, Group, TextInput } from "@mantine/core";
import { Menu } from "../icons/Menu";
import { Glass } from "../icons/Glass";
import { Location } from "../Location";
import { DeliveryTime } from "../DeliveryTime";
import { User } from "../icons/User";
import styles from "./Header.module.css";
import { CartButton } from "./components/CartButton";

type Props = {
  onNavbar: () => void;
};

export const Header = ({ onNavbar }: Props) => {
  return (
    <Flex
      mx="auto"
      maw={1454}
      className={styles.root}
      align="center"
      mah={82}
      px={20}
      py={20}
      justify="space-between"
    >
      <Group w="100%" wrap="nowrap">
        <Button onClick={onNavbar} h={38} w={38} px={8} variant="secondary">
          <Menu size={24} />
        </Button>

        <TextInput
          w="100%"
          maw={400}
          size="md"
          leftSection={<Glass size={16} />}
          placeholder="Поиск товаров"
        />

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
  );
};
