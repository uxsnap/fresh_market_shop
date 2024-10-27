import { Trash } from "@/components/icons/Trash";
import { useCartStore } from "@/store";
import { Button } from "@mantine/core";
import { MouseEvent } from "react";

import styles from "./RemoveAll.module.css";

export const RemoveAll = () => {
  const removeAllItems = useCartStore((state) => state.removeAllItems);

  const handleRemoveAll = (e: MouseEvent<HTMLElement>) => {
    e.stopPropagation();
    removeAllItems();
  };

  return (
    <Button
      className={styles.root}
      variant="accent-reverse"
      leftSection={<Trash size={20} fill="var(--mantine-color-accent-0)" />}
      onClick={handleRemoveAll}
    >
      Очистить корзину
    </Button>
  );
};
