import { Group, Text } from "@mantine/core";
import cn from "classnames";
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
import { MainIcon } from "../icons/MainIcon";

import styles from "./CategoryItem.module.css";
import { MouseEventHandler } from "react";

const mapNameToIcon: Record<string, React.FC> = {
  Главная: MainIcon,
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
  onClick: MouseEventHandler<HTMLDivElement>;
  active?: boolean;
};

export const CategoryItem = ({ children, onClick, active }: Props) => {
  const Icon = mapNameToIcon[children];

  return (
    <Group
      align="center"
      py={4}
      px={12}
      className={cn(styles.item, active && styles.active)}
      key={children}
      gap={10}
      onClick={onClick}
    >
      {Icon ? <Icon /> : ""}

      <Text lh={1} fw={500} fz={18} c="accent.0">
        {children}
      </Text>
    </Group>
  );
};
