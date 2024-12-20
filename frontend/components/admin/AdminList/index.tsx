import { getAdmins } from "@/api/auth/getAdmins";
import { Table } from "@mantine/core";
import { useQuery } from "@tanstack/react-query";
import { memo } from "react";

export const AdminList = memo(() => {
  const { data, isFetched } = useQuery({
    queryFn: getAdmins,
    queryKey: [getAdmins.queryKey],
    refetchOnWindowFocus: false,
    staleTime: Infinity,
  });

  console.log(data?.data);

  const rows = data?.data.admins.map((admin) => (
    <Table.Tr key={admin.name}>
      <Table.Td>{admin.uid}</Table.Td>
      <Table.Td>{admin.email}</Table.Td>
    </Table.Tr>
  ));

  return (
    <Table>
      <Table.Thead>
        <Table.Tr>
          <Table.Th>ID</Table.Th>
          <Table.Th>Mail</Table.Th>
        </Table.Tr>
      </Table.Thead>

      {isFetched && <Table.Tbody>{rows}</Table.Tbody>}
    </Table>
  );
});
