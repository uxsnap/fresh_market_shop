import { Group, Image, Modal, Stack, Title } from "@mantine/core";

import styles from "./RecipeModal.module.css";
import { ItemList } from "../ItemList";
import { RecipeStep } from "../RecipeStep";
import { RecipeProducts } from "./RecipeProducts";

type Props = {
  close: () => void;
  uid: string;
  name: string;
  ccal: number;
  img: string;
};

export const RecipeModal = ({
  close,
  uid,
  name = "Название рецепта",
  ccal = 300,
  img = "/recipe.png",
}: Props) => {
  return (
    <Modal.Root zIndex={100} opened={true} onClose={close}>
      <Modal.Overlay />

      <Modal.Content maw={640} miw={640}>
        <Modal.Header className={styles.header} px={20} py={12}>
          <Group w="100%" gap={16} wrap="nowrap">
            <Title order={2} lineClamp={1} c="accent.0">
              {name}
            </Title>

            <Title textWrap="nowrap" order={4} c="accent.2">
              {ccal} ккал
            </Title>
          </Group>

          <Modal.CloseButton size="32px" c="accent.0" />
        </Modal.Header>

        <Modal.Body p={0}>
          <Image mah={194} src={img} />

          <Stack p={12} style={{ overflowX: "auto", overflowY: "hidden" }}>
            <RecipeProducts uid={uid} />

            {Array.from({ length: 5 }).map((_, ind) => (
              <RecipeStep key={ind} step={ind + 1} maxStep={5} />
            ))}
          </Stack>
        </Modal.Body>
      </Modal.Content>
    </Modal.Root>
  );
};
