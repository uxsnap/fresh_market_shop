"use client";

import { Header } from "@/components/Header";
import { Recipe } from "@/components/Recipe";
import { RecipeModal } from "@/components/RecipeModal";
import { SideMenu } from "@/components/SideMenu";
import { AppShell, Container } from "@mantine/core";

export default function HomePage() {
  return (
    <Container p={10}>
      <RecipeModal />
    </Container>
    // <AppShell
    //   header={{ height: 78 }}
    //   navbar={{
    //     width: 300,
    //     breakpoint: "sm",
    //   }}
    //   padding="md"
    // >
    //   <AppShell.Header>
    //       <Header />
    //   </AppShell.Header>

    //   <AppShell.Navbar px={12} py={20}>
    //     <SideMenu />
    //   </AppShell.Navbar>

    //   <AppShell.Main>Main</AppShell.Main>
    // </AppShell>
  );
}
