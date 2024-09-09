import { Container, Flex, Stack, Title, Text } from "@mantine/core";
import { Slider } from "../Slider/Slider";
import { Salad } from "../icons/Salad";
import { Bread } from "../icons/Bread";
import { Fish } from "../icons/Fish";
import { Peach } from "../icons/Peach";
import { Tomato } from "../icons/Tomato";
import { Meat } from "../icons/Meat";
import { Oil } from "../icons/Oil";
import { Milk } from "../icons/Milk";
import { Grains } from "../icons/Grains";
import { Bean } from "../icons/Bean";

import styles from "./SideMenu.module.css";

const menu = [
  { label: "Готовая еда", icon: Salad },
  { label: "Хлебобулочные изделия", icon: Bread },
  { label: "Рыба", icon: Fish },
  { label: "Фрукты", icon: Peach },
  { label: "Овощи", icon: Tomato },
  { label: "Мясной отдел", icon: Meat },
  { label: "Бакалея", icon: Grains },
  { label: "Заправка", icon: Oil },
  { label: "Молочные продукты", icon: Milk },
  { label: "Топпинги", icon: Bean },
];

export const SideMenu = () => {
  return (
    <Container>
      <Stack gap={20}>
        <Title c="accent.0" order={2}>
          Каталог
        </Title>

        <Stack>
          <Title order={4} c="accent.0">
            Кол-во каллорий:
          </Title>
          <Slider />
        </Stack>

        <Stack mt={20} gap={12}>
          {menu.map(({ label, icon: Icon }) => (
            <Flex py={4} px={12} className={styles.item} key={label} gap={10}>
              <Icon />

              <Text c="accent.0">{label}</Text>
            </Flex>
          ))}
        </Stack>
      </Stack>
    </Container>
  );
};
