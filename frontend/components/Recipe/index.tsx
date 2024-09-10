import { Text, Title, Stack, BackgroundImage, Box } from "@mantine/core";

import styles from "./Recipe.module.css";

export const Recipe = () => {
  return (
    <BackgroundImage
      mih={280}
      maw={200}
      display="flex"
      src="https://raw.githubusercontent.com/mantinedev/mantine/master/.demo/images/bg-8.png"
      radius="lg"
      pos="relative"
      className={styles.root}
    >
      <Stack w="100%" className={styles.main} gap={16} justify="flex-end">
        <Box py={20} px={16}>
          <Title order={3} c="accent.0">
            Название Рецепта
          </Title>
          <Text fw={500} fz={18} c="accent.1">
            1 час 20 минут
          </Text>
        </Box>
      </Stack>
    </BackgroundImage>
  );
};
