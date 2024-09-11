import { Group, Image, Modal, Stack, Title } from "@mantine/core";

import styles from "./RecipeModal.module.css";
import { ItemList } from "../ItemList";
import { RecipeStep } from "../RecipeStep";

type Props = {
  name?: string;
  ccal?: string;
};

export const RecipeModal = ({
  name = "Название рецепта",
  ccal = "300 ккал",
}: Props) => {
  return (
    <Modal.Root opened={true} onClose={close}>
      <Modal.Overlay />

      <Modal.Content maw={640} miw={640}>
        <Modal.Header className={styles.header} px={20} py={12}>
          <Group gap={16}>
            <Title c="accent.0">{name}</Title>

            <Title order={4} c="accent.2">
              {ccal}
            </Title>
          </Group>

          <Modal.CloseButton size="32px" c="accent.0" />
        </Modal.Header>

        <Modal.Body p={0}>
          <Image mah={194} src="/recipe.png" />

          <Stack p={12} style={{ overflowX: "auto", overflowY: "hidden" }}>
            <ItemList noTitle type="small" />

            {Array.from({ length: 5 }).map((_, ind) => (
              <RecipeStep key={ind} step={ind + 1} maxStep={5} />
            ))}
          </Stack>
        </Modal.Body>
      </Modal.Content>
    </Modal.Root>
  );
};
