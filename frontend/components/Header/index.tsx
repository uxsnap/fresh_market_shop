import { Button, Flex, Group, TextInput } from "@mantine/core";
import { Menu } from "../icons/Menu";
import { Glass } from "../icons/Glass";
import { Location } from "../Location";

export const Header = () => {
  return (
    <Flex align="stretch" mah={82} px={20} py={20} justify="space-between">
      <Group>
        <Button h={40} w={40} px={8} variant="secondary">
          <Menu />
        </Button>

        <TextInput
          size="md"
          leftSection={<Glass size={16} />}
          placeholder="Поиск товаров"
        />

        <Location />
      </Group>

      <Group></Group>
    </Flex>
  );
};
