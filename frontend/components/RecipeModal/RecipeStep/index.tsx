import { Group, Image, Stack, Title, Text } from "@mantine/core";

import styles from "./RecipeStep.module.css";

type Props = {
  step: number;
  maxStep: number;
  children?: string;
  src?: string;
};

export const RecipeStep = ({
  src = "recipe.png",
  step,
  maxStep,
  children,
}: Props) => {
  return (
    <Group
      wrap="nowrap"
      mah={180}
      className={styles.root}
      align="flex-start"
      gap={20}
      w="100%"
    >
      <Image className={styles.img} src={src} />

      <Stack style={{ overflow: 'auto' }} className={styles.info}>
        <Title order={4} c="accent.2">
          Шаг {step} из {maxStep}
        </Title>

        <Text style={{ overflowY: "auto" }} fz={14} c="accent.0">
          {children}
        </Text>
      </Stack>
    </Group>
  );
};
