"use client";

import { AdminList } from "@/components/admin/AdminList";
import { AdminProductList } from "@/components/admin/AdminProductList";
import { AdminRecipeList } from "@/components/admin/AdminRecipeList";
import { CreateButton } from "@/components/admin/CreateButton";
import { useAdminStore } from "@/store/admin";
import { AdminTab } from "@/types";
import { Box, Group, Stack, Tabs } from "@mantine/core";

export default function AdminPage() {
  const tab = useAdminStore((s) => s.tab);
  const setTab = useAdminStore((s) => s.setTab);

  return (
    <Stack gap={24} p={8} m={0} miw="100%">
      <Tabs
        value={tab}
        onChange={(v) => setTab(v as AdminTab)}
        color="accent.0"
        variant="pills"
        defaultValue="admins"
      >
        <Group justify="space-between">
          <Tabs.List>
            <Tabs.Tab value="admins">Администраторы</Tabs.Tab>
            <Tabs.Tab value="products">Продукты</Tabs.Tab>
            <Tabs.Tab value="recipes">Рецепты</Tabs.Tab>
          </Tabs.List>

          <CreateButton />
        </Group>

        <Box mt={24}>
          <Tabs.Panel value="admins">
            <AdminList />
          </Tabs.Panel>

          <Tabs.Panel value="products">
            <AdminProductList />
          </Tabs.Panel>

          <Tabs.Panel value="recipes">
            <AdminRecipeList />
          </Tabs.Panel>
        </Box>
      </Tabs>
    </Stack>
  );
}
