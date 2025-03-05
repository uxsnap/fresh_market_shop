"use client";

import { AdminList } from "@/components/admin/AdminList";
import { AdminProductList } from "@/components/admin/AdminProductList";
import { AdminRecipeList } from "@/components/admin/AdminRecipeList";
import { CreateButton } from "@/components/admin/CreateButton";
import { AdminTab } from "@/types";
import { Box, Group, Stack, Tabs } from "@mantine/core";
import { usePathname, useRouter, useSearchParams } from "next/navigation";
import { useEffect, useMemo } from "react";

export default function AdminPage() {
  const pathname = usePathname();
  const router = useRouter();
  const searchParams = useSearchParams();
  const params = useMemo(() => new URLSearchParams(), []);

  const tab = (searchParams.get("tab") ?? AdminTab.admins) as AdminTab;

  const setSearchParams = (v: AdminTab) => {
    params.set("tab", v);

    router.push(`${pathname}?${params.toString()}`);
  };

  useEffect(() => {
    setSearchParams(tab);
  }, []);

  return (
    <Stack gap={24} p={8} m={0} miw="100%">
      <Tabs
        value={tab}
        onChange={(v) => setSearchParams(v as AdminTab)}
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
