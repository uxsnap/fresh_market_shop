import { Group, Text } from "@mantine/core";

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

import styles from './CategoryItem.module.css'

const mapNameToIcon: Record<string, React.FC> = {
  "Готовая еда": Salad,
  Хлеб: Bread,
  Рыба: Fish,
  Фрукты: Peach,
  Овощи: Tomato,
  "Мясной отдел": Meat,
  Бакалея: Grains,
  Заправка: Oil,
  "Молочные продукты": Milk,
  Топпинги: Bean,
};

type Props = {
  children: string;
};

export const CategoryItem = ({ children }: Props) => {
  const Icon = mapNameToIcon[children];

  return (
    <Group
      align="center"
      py={4}
      px={12}
      className={styles.item}
      key={children}
      gap={10}
    >
      {Icon ? <Icon /> : ""}

      <Text lh={1} fw={500} fz={18} c="accent.0">
        {children}
      </Text>
    </Group>
  );
};
