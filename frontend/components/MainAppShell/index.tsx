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

  const [mobileOpened, { toggle: toggleMobile, close: closeMobile }] =
    useDisclosure();
  const [desktopOpened, { toggle: toggleDesktop, close: closeDesktop }] =
    useDisclosure(true);

  useEffect(() => {
    if (pathname === "/profile") {
      closeMobile();
      closeDesktop();
    }
  }, [pathname]);

  const handleToggle = () => {
    toggleDesktop();
    toggleMobile();
  };

  return (
    <QueryClientProvider client={queryClient}>
      <AppShell
        header={{ height: { base: 125, md: 78 } }}
        navbar={{
          width: 300,
          breakpoint: "md",
          collapsed: { mobile: !mobileOpened, desktop: !desktopOpened },
        }}
        padding="md"
      >
        <AppShell.Header zIndex={3}>
          <Header onNavbar={handleToggle} />
        </AppShell.Header>

        <AppShell.Navbar zIndex={1} px={12} py={20}>
          <SideMenu />
        </AppShell.Navbar>

        <AppShell.Main>{children}</AppShell.Main>
      </AppShell>
    </QueryClientProvider>
  );
};
