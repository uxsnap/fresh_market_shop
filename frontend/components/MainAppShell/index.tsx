"use client";

import { AppShell, useMatches } from "@mantine/core";
import { Header } from "../Header";
import { SideMenu } from "../SideMenu";
import { PropsWithChildren, useEffect } from "react";
import { usePathname, useRouter } from "next/navigation";
import { useDisclosure } from "@mantine/hooks";
import {
  QueryClient,
  QueryClientProvider,
  useMutation,
} from "@tanstack/react-query";
import { verifyUser } from "@/api/auth/verify";
import { useAuthStore } from "@/store/auth";
import { ItemCardExtended } from "../ItemCard/ItemExtendedCard";

const queryClient = new QueryClient();

const MainApp = ({ children }: PropsWithChildren) => {
  const pathname = usePathname();
  const router = useRouter();
  const setLogged = useAuthStore((s) => s.setLogged);

  const isDesktop = useMatches({
    base: false,
    md: true,
  });

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

  const { mutate } = useMutation({
    mutationFn: verifyUser,
    onSuccess: ({ isValid }) => {
      if (!isValid) {
        router.push("/");
      }

      setLogged(isValid);
    },
  });

  useEffect(() => {
    mutate();
  }, [children]);

  return (
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
        <Header
          opened={isDesktop ? desktopOpened : mobileOpened}
          onNavbar={handleToggle}
        />
      </AppShell.Header>

      <AppShell.Navbar zIndex={1} px={12} py={20}>
        <SideMenu />
      </AppShell.Navbar>

      <AppShell.Main>
        {children}

        <ItemCardExtended />
      </AppShell.Main>
    </AppShell>
  );
};

export const MainAppShell = ({ children }: PropsWithChildren) => (
  <QueryClientProvider client={queryClient}>
    <MainApp>{children}</MainApp>
  </QueryClientProvider>
);
