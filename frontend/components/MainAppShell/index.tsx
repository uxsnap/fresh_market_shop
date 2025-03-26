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
import { Support } from "../Support";
import { AdminHeader } from "../admin/Header";

const queryClient = new QueryClient();

const MainApp = ({ children }: PropsWithChildren) => {
  const pathname = usePathname();
  const router = useRouter();
  const setLogged = useAuthStore((s) => s.setLogged);
  const setAdmin = useAuthStore((s) => s.setAdmin);
  const admin = useAuthStore((s) => s.admin);

  const isAdmin = admin && pathname.startsWith("/admin");

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

  const { mutate, isSuccess } = useMutation({
    mutationFn: verifyUser,
    onSuccess: ({ isValid, isAdmin }) => {
      if (!isValid) {
        router.push("/");
      }

      setLogged(isValid);
      setAdmin(isAdmin);
    },
  });

  useEffect(() => {
    mutate();
  }, []);

  return (
    <AppShell
      header={{ height: { base: 178, md: 78 } }}
      navbar={{
        width: isAdmin ? 0 : 300,
        breakpoint: "md",
        collapsed: { mobile: !mobileOpened, desktop: !desktopOpened },
      }}
      padding="md"
    >
      <AppShell.Header zIndex={3}>
        {isAdmin ? (
          <AdminHeader />
        ) : (
          <Header
            opened={isDesktop ? desktopOpened : mobileOpened}
            onNavbar={handleToggle}
          />
        )}
      </AppShell.Header>

      {!isAdmin && (
        <AppShell.Navbar zIndex={1} px={12} py={20}>
          <SideMenu onNavbar={handleToggle} />
        </AppShell.Navbar>
      )}

      <AppShell.Main>
        {children}

        <ItemCardExtended />

        {isSuccess && !isAdmin && <Support />}
      </AppShell.Main>
    </AppShell>
  );
};

export const MainAppShell = ({ children }: PropsWithChildren) => (
  <QueryClientProvider client={queryClient}>
    <MainApp>{children}</MainApp>
  </QueryClientProvider>
);
