import { Group, Image, Stack, Title, Text } from "@mantine/core";

import styles from "./RecipeStep.module.css";

type Props = {
  step: number;
  maxStep: number;
  children?: string;
  src?: string;
};

export const RecipeStep = ({ src = "recipe.png", step, maxStep }: Props) => {
  return (
    <Group
      wrap="nowrap"
      align="flex-start"
      mah={180}
      className={styles.root}
      gap={20}
    >
      <Image className={styles.img} w={200} mah={180} src={src} />

      <Stack mah={180} py={16} gap={12}>
        <Title order={4} c="accent.2">
          Шаг {step} из {maxStep}
        </Title>

        <Text style={{ overflowY: "auto" }} fz={14} c="accent.0">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
          eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad
          minim veniam, quis nostrud exercitation ullamco laboris nisi ut
          aliquip ex ea commodo consequat. Duis aute irure dolor in
          reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
          pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
          culpa qui officia deserunt mollit anim id est laborum.
        </Text>
      </Stack>
    </Group>
  );
};
