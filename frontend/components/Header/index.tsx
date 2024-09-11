import { Button, Flex, Group, TextInput } from "@mantine/core";
import { Menu } from "../icons/Menu";
import { Glass } from "../icons/Glass";
import { Location } from "../Location";
import { DeliveryTime } from "../DeliveryTime";
import { User } from "../icons/User";
import styles from "./Header.module.css";
import { CartButton } from "./components/CartButton";

export const Header = () => {
  return (
    <Flex mx="auto" maw={1454} className={styles.root} align="center" mah={82} px={20} py={20} justify="space-between">
      <Group>
        <Button h={38} w={38} px={8} variant="secondary">
          <Menu />
        </Button>

        <TextInput
          miw={400}
          size="md"
          leftSection={<Glass size={16} />}
          placeholder="Поиск товаров"
        />

        <Location />

        <DeliveryTime />
      </Group>

      <Group gap={24} align="center">
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
