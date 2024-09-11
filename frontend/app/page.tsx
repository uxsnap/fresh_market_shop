"use client";

import { Header } from "@/components/Header";
import { SideMenu } from "@/components/SideMenu";
import { AppShell, Flex } from "@mantine/core";

export default function HomePage() {
  return (
    <AppShell
      header={{ height: 78 }}
      navbar={{
        width: 300,
        breakpoint: "sm",
      }}
      padding="md"
    >
      <AppShell.Header>
          <Header />
      </AppShell.Header>

      <AppShell.Navbar px={12} py={20}>
        <SideMenu />
      </AppShell.Navbar>

      <AppShell.Main>Main</AppShell.Main>
    </AppShell>
  );
}
