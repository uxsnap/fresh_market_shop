import { Delivery } from "@/components/icons/Delivery";
import { useCartStore } from "@/store";
import { Box, Button, Group, Text } from "@mantine/core";

import styles from "./PayButton.module.css";

export const PayButton = () => {
  const price = useCartStore((s) => s.getFullPrice());

  if (!price) {
    return null;
  }

  return (
    <Box hiddenFrom="md" className={styles.root}>
      <Button w="100%" variant="accent" h={40}>
        <Group gap={16} align="center">
          <Text fw="bold" fz={18}>
            Оформить заказ
          </Text>

          <Group gap={8} align="center">
            <Text fw="bold" fz={18}>
              {price} ₽
            </Text>
            <Delivery size={24} fill="var(--mantine-color-accent-0)" />
          </Group>
        </Group>
      </Button>
    </Box>
  );
};
