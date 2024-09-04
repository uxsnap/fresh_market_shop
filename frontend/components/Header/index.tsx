import { Button, Flex, Group, TextInput } from "@mantine/core";
import { Menu } from "../icons/Menu";
import { Glass } from "../icons/Glass";
import { Location } from "../Location";
import { DeliveryTime } from "../DeliveryTime";

export const Header = () => {
  return (
    <Flex align="stretch" mah={82} px={20} py={20} justify="space-between">
      <Group>
        <Button h={38} w={38} px={8} variant="secondary">
          <Menu />
        </Button>

        <TextInput
          size="md"
          leftSection={<Glass size={16} />}
          placeholder="Поиск товаров"
        />

        <Location />

        <DeliveryTime />
      </Group>

      <Group></Group>
    </Flex>
  );
};
