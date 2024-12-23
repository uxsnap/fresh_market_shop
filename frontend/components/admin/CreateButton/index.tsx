import { useAdminStore } from "@/store/admin";
import { AdminTab } from "@/types";
import { Button, Group, Modal, Stack, Title } from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import { useMemo } from "react";
import { CreateAdminModal } from "./CreateAdminModal";

const mapTabToText: Record<AdminTab, string> = {
  [AdminTab.admins]: "Создать админа",
};

export const CreateButton = () => {
  const tab = useAdminStore((s) => s.tab);

  const [opened, { open, close }] = useDisclosure(false);

  const text = useMemo(() => mapTabToText[tab], [tab]);

  return (
    <>
      <Modal
        centered
        opened={opened}
        onClose={close}
        title={
          <Title c="accent.0" order={3}>
            {text}
          </Title>
        }
      >
        <CreateAdminModal onClose={close} />
      </Modal>

      <Button variant="accent" onClick={open}>
        {text}
      </Button>
    </>
  );
};
