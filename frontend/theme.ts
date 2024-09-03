"use client";

import {
  createTheme,
  defaultVariantColorsResolver,
  parseThemeColor,
} from "@mantine/core";

export const theme = createTheme({
  colors: {
    accent: ["#4F463D", "#706962", "", "", "", "", "", "", "", ""],
    secondary: ["#FEBDB9", "#FFA49E", "", "", "", "", "", "", "", ""],
    primary: ["#578C3E", "#B7D968", "#DCEABD", "", "", "", "", "", "", ""],
    bg: ["#F0EEEE", "#F8F8F8", "#FFFFFF", "", "", "", "", "", "", ""],
    danger: ["#B50000", "#8A0606", "", "", "", "", "", "", "", ""],
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
  variantColorResolver: (input) => {
    const defaultResolvedColors = defaultVariantColorsResolver(input);

    switch (input.variant) {
      case "accent":
        return {
          ...defaultResolvedColors,
          color: "var(--mantine-color-white)",
          background: "var(--mantine-color-accent-0)",
          hover: "var(--mantine-color-accent-1)",
        };
      case "accent-reverse":
        return {
          ...defaultResolvedColors,
          color: "var(--mantine-color-accent-0)",
          border: "1px solid var(--mantine-color-accent-0)",
          background: "var(--mantine-color-bg-2)",
          hover: "var(--mantine-color-bg-0)",
        };
      case "secondary":
        return {
          ...defaultResolvedColors,
          color: "var(--mantine-color-accent-0)",
          border: "1px solid var(--mantine-color-accent-0)",
          background: "var(--mantine-color-secondary-0)",
          hover: "var(--mantine-color-secondary-1)",
        };
      case "outline":
        return {
          color: "var(--mantine-color-accent-0)",
          border: "",
          background: "transparent",
          hover: "transparent",
        };
      case "danger":
        return {
          ...defaultResolvedColors,
          color: "var(--mantine-color-white)",
          background: "var(--mantine-color-danger-0)",
          hover: "var(--mantine-color-danger-1)",
        };
      case "dashed":
        return {
          ...defaultResolvedColors,
          color: "var(--mantine-color-accent-0)",
          border: "1px dashed var(--mantine-color-accent-0)",
          background: "transparent",
          hover: "transparent",
        };
      default:
        return defaultResolvedColors;
    }
  },
});
