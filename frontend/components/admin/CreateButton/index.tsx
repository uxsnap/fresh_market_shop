import { useAdminStore } from "@/store/admin";
import { AdminTab } from "@/types";
import { Button, Modal, Title } from "@mantine/core";
import { useEffect, useMemo } from "react";
import { CreateAdminModal } from "./CreateAdminModal";
import { ProductModal } from "./ProductModal";
import { RecipeModal } from "./components/RecipeModal";
import { useSearchParams } from "next/navigation";

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
  const searchParams = useSearchParams();

  const tab = (searchParams.get("tab") ?? AdminTab.admins) as AdminTab;

  const modalOpen = useAdminStore((s) => s.modalOpen);
  const setModalOpen = useAdminStore((s) => s.setModalOpen);
  const setRecipeItem = useAdminStore((s) => s.setRecipeItem);

  const text = useMemo(() => mapTabToText[tab], [tab]);
  const ModalContents = useMemo(() => mapTabToModal[tab], [tab]);

  const handleClose = () => {
    setModalOpen(false);
    setRecipeItem(undefined);
  };

  return (
    <>
      <Modal
        centered
        opened={!!modalOpen}
        onClose={() => handleClose()}
        title={
          <Title c="accent.0" order={3}>
            {text}
          </Title>
        }
      >
        <ModalContents onClose={() => handleClose()} />
      </Modal>

      <Button variant="accent" onClick={() => setModalOpen(true)}>
        {text}
      </Button>
    </>
  );
};
