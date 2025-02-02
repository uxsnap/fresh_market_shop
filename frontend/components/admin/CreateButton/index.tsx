import { useAdminStore } from "@/store/admin";
import { AdminTab } from "@/types";
import { Button, Modal, Title } from "@mantine/core";
import { useMemo } from "react";
import { CreateAdminModal } from "./CreateAdminModal";
import { ProductModal } from "./ProductModal";
import { RecipeModal } from "./RecipeModal";

const mapTabToText: Record<AdminTab, string> = {
  [AdminTab.admins]: "Создать админа",
  [AdminTab.products]: "Добавить продукт",
  [AdminTab.recipes]: "Добавить рецепт",
};

const mapTabToModal: Record<
  AdminTab,
  ({ onClose }: { onClose: () => void }) => JSX.Element
> = {
  [AdminTab.admins]: CreateAdminModal,
  [AdminTab.products]: ProductModal,
  [AdminTab.recipes]: RecipeModal,
};

export const CreateButton = () => {
  const tab = useAdminStore((s) => s.tab);
  const modalOpen = useAdminStore((s) => s.modalOpen);
  const setModalOpen = useAdminStore((s) => s.setModalOpen);

  const text = useMemo(() => mapTabToText[tab], [tab]);
  const ModalContents = useMemo(() => mapTabToModal[tab], [tab]);

  return (
    <>
      <Modal
        centered
        opened={!!modalOpen}
        onClose={() => setModalOpen(false)}
        title={
          <Title c="accent.0" order={3}>
            {text}
          </Title>
        }
      >
        <ModalContents onClose={() => setModalOpen(false)} />
      </Modal>

      <Button variant="accent" onClick={() => setModalOpen(true)}>
        {text}
      </Button>
    </>
  );
};
