import "@mantine/core/styles.css";
import "@mantine/carousel/styles.css";
import "@mantine/notifications/styles.css";
import "@mantine/dropzone/styles.css";

import React from "react";
import { ColorSchemeScript, MantineProvider } from "@mantine/core";
import { theme } from "../theme";
import { MainAppShell } from "@/components/MainAppShell";
import { Notifications } from "@mantine/notifications";

export const metadata = {
  title: "Fresh Market Shop",
};

export default function RootLayout({ children }: { children: any }) {
  return (
    <html lang="en">
      <head>
        <ColorSchemeScript />
        <link rel="shortcut icon" href="/favicon.svg" />
        <meta
          name="viewport"
          content="minimum-scale=1, initial-scale=1, width=device-width, user-scalable=no"
        />
      </head>
      <body>
        <MantineProvider theme={theme}>
          <Notifications />

          <MainAppShell>{children}</MainAppShell>
        </MantineProvider>
      </body>
    </html>
  );
}
