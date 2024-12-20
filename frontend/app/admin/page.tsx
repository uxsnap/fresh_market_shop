"use client";

import { AdminList } from "@/components/admin/AdminList";
import { Stack, Tabs } from "@mantine/core";

export default function AdminPage() {
  return (
    <Stack gap={24} p={8} m={0} miw="100%">
      <Tabs color="accent.0" variant="pills" defaultValue="gallery">
        <Tabs.List>
          <Tabs.Tab value="admins">Администраторы</Tabs.Tab>
        </Tabs.List>

        <Tabs.Panel value="admins">
          <AdminList />
        </Tabs.Panel>
      </Tabs>
    </Stack>
  );
}
