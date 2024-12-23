import { deleteAccount } from "@/api/auth/deleteAccount";
import { getAdmins } from "@/api/auth/getAdmins";
import { Close } from "@/components/icons/Close";
import { Admin } from "@/types";
import {
  showErrorNotification,
  showInlineErrorNotification,
  showSuccessNotification,
} from "@/utils";
import {
  ActionIcon,
  Box,
  Button,
  Group,
  Modal,
  Table,
  Title,
} from "@mantine/core";
import { useMutation, useQuery } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { memo, useEffect, useState } from "react";

export const AdminList = memo(() => {
  const [deleteCandidate, setDeleteCandidate] = useState<Admin | undefined>();

  const { data, isFetched, refetch, error } = useQuery({
    queryFn: getAdmins,
    queryKey: [getAdmins.queryKey],
    refetchOnWindowFocus: false,
    staleTime: Infinity,
    retry: false,
  });

  useEffect(() => {
    if (error) {
      showErrorNotification(error as AxiosError<any>);
    }
  }, [error]);

  const { mutate: mutateDelete, isPending: isPendingDelete } = useMutation({
    mutationFn: deleteAccount,
    mutationKey: [deleteAccount.queryKey],
    onError: () => {
      showInlineErrorNotification("Ошибка удаления пользователя!");
      setDeleteCandidate(undefined);
    },
    onSuccess: () => {
      showSuccessNotification("Пользователь был удален!");
      refetch();
      setDeleteCandidate(undefined);
    },
  });

  const rows = data?.data.admins.map((admin) => (
    <Table.Tr key={admin.uid}>
      <Table.Td>{admin.uid}</Table.Td>
      <Table.Td>{admin.email}</Table.Td>

      <Table.Td>
        <ActionIcon
          onClick={() => setDeleteCandidate(admin)}
          variant="transparent"
          aria-label="Close"
        >
          <Close fill="var(--mantine-color-danger-2)" />
        </ActionIcon>
      </Table.Td>
    </Table.Tr>
  ));

  return (
    <Box>
      <Modal
        opened={!!deleteCandidate}
        onClose={close}
        title={
          <Title c="accent.0" order={3}>
            Удалить пользователя {deleteCandidate?.email}?
          </Title>
        }
        centered
      >
        <Group wrap="nowrap">
          <Button
            variant="danger"
            onClick={() => mutateDelete({ uid: deleteCandidate!.uid })}
            disabled={isPendingDelete}
            w="100%"
            mih={32}
          >
            Удалить
          </Button>

          <Button
            variant="accent-reverse"
            onClick={() => setDeleteCandidate(undefined)}
            w="100%"
            mih={32}
          >
            Закрыть
          </Button>
        </Group>
      </Modal>

      <Table>
        <Table.Thead>
          <Table.Tr>
            <Table.Th>ID</Table.Th>
            <Table.Th>Mail</Table.Th>
            <Table.Th>Действия</Table.Th>
          </Table.Tr>
        </Table.Thead>

        {isFetched && <Table.Tbody>{rows}</Table.Tbody>}
      </Table>
    </Box>
  );
});
