"use client";

import { AppShell } from "@mantine/core";
import { Header } from "../Header";
import { SideMenu } from "../SideMenu";
import { PropsWithChildren, useEffect } from "react";
import { usePathname } from "next/navigation";
import { useDisclosure } from "@mantine/hooks";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

const queryClient = new QueryClient();

export const MainAppShell = ({ children }: PropsWithChildren) => {
  const pathname = usePathname();

  const [desktopOpened, { close, toggle: toggleDesktop }] = useDisclosure(true);

  useEffect(() => {
    if (pathname === "/profile") {
      close();
    }
  }, [pathname]);

  return (
    <QueryClientProvider client={queryClient}>
      <AppShell
        header={{ height: 78 }}
        navbar={{
          width: 300,
          breakpoint: "sm",
          collapsed: { desktop: !desktopOpened },
        }}
        padding="md"
      >
        <AppShell.Header>
          <Header onNavbar={toggleDesktop} />
        </AppShell.Header>

        <AppShell.Navbar px={12} py={20}>
          <SideMenu />
        </AppShell.Navbar>

        <AppShell.Main>{children}</AppShell.Main>
      </AppShell>
    </QueryClientProvider>
  );
};
