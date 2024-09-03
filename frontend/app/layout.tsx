import "@mantine/core/styles.css";

import React from "react";
import { ColorSchemeScript, MantineProvider } from "@mantine/core";
import { theme } from "../theme";
import { Roboto } from "next/font/google";

export const metadata = {
  title: "Fresh Market Shop",
};

const roboto = Roboto({
  subsets: ["cyrillic"],
  weight: ["400", "500", "700"],
});

export default function RootLayout({ children }: { children: any }) {
  return (
    <html lang="en" className={roboto.className}>
      <head>
        <ColorSchemeScript />
        <link rel="shortcut icon" href="/favicon.svg" />
        <meta
          name="viewport"
          content="minimum-scale=1, initial-scale=1, width=device-width, user-scalable=no"
        />
      </head>
      <body>
        <MantineProvider theme={theme}>{children}</MantineProvider>
      </body>
    </html>
  );
}
