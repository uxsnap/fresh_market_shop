"use client";

import { createTheme } from "@mantine/core";

export const theme = createTheme({
  colors: {
    accent: ["#4F463D", "#706962", "", "", "", "", "", "", "", ""],
    secondary: ["#FEBDB9", "#FFA49E", "", "", "", "", "", "", "", ""],
    primary: ["#578C3E", "#B7D968", "#DCEABD", "", "", "", "", "", "", ""],
    bg: ["#F0EEEE", "#F8F8F8", "#FFFFFF", "", "", "", "", "", "", ""],
    error: ["#B50000", "", "", "", "", "", "", "", "", ""],
  },
  defaultRadius: "md",
  radius: {
    md: "8px",
  },
  headings: {
    sizes: {
      h1: {
        fontSize: "32px",
      },
      h2: {
        fontSize: "26px",
      },
      h3: {
        fontSize: "22px",
      },
      h4: {
        fontSize: "18px",
      },
      h5: {
        fontSize: "14px",
      },
      h6: {
        fontSize: "12px",
      },
    },
  },
});
