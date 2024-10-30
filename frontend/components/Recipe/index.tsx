import { Text, Title, Stack, BackgroundImage, Box } from "@mantine/core";

import styles from "./Recipe.module.css";

export type Props = {
  name: string;
  time: string;
  onClick?: () => void;
  img?: string;
};

export const Recipe = ({ onClick, name, time, img }: Props) => {
  const fallbackImg =
    "https://raw.githubusercontent.com/mantinedev/mantine/master/.demo/images/bg-8.png";

  return (
    <BackgroundImage
      mih={280}
      w={200}
      display="flex"
      src={img ?? fallbackImg}
      radius="lg"
      pos="relative"
      className={styles.root}
      onClick={onClick}
    >
      <Stack w="100%" className={styles.main} gap={16} justify="flex-end">
        <Box py={20} px={16}>
          <Title order={3} c="accent.0">
            {name}
          </Title>
          <Text fw={500} fz={18} c="accent.1">
            {time}
          </Text>
        </Box>
      </Stack>
    </BackgroundImage>
  );
};
