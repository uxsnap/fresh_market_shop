"use client";

import {
  Button,
  createTheme,
  DEFAULT_THEME,
  defaultVariantColorsResolver,
  FileInput,
  Input,
  PasswordInput,
  RangeSlider,
  rem,
  Slider,
  TextInput,
} from "@mantine/core";

import classes from "./Main.module.css";
import { robotoFont } from "./static/roboto/Roboto";

export const theme = createTheme({
  colors: {
    accent: ["#4F463D", "#706962", "#C3C0BD", "", "", "", "", "", "", ""],
    secondary: ["#FEBDB9", "#FFA49E", "", "", "", "", "", "", "", ""],
    primary: ["#578C3E", "#B7D968", "#DCEABD", "", "", "", "", "", "", ""],
    bg: ["#F0EEEE", "#F8F8F8", "#FFFFFF", "", "", "", "", "", "", ""],
    danger: ["#B50000", "#8A0606", "#C23131", "", "", "", "", "", "", ""],
    warning: ["#FFBE00", "", "", "", "", "", "", "", "", ""],
  },
  defaultRadius: "md",
  radius: {
    md: "8px",
    lg: "12px",
  },
  fontFamily: robotoFont.style.fontFamily,
  fontFamilyMonospace: "Monaco, Courier, monospace",
  headings: {
    fontFamily: `${robotoFont.style.fontFamily}, ${DEFAULT_THEME.fontFamily}`,

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
  components: {
    TextInput: TextInput.extend({
      vars: (_, props) => {
        switch (props.size) {
          case "md":
            return {
              wrapper: {
                "--input-height": rem(42),
                "--input-padding-x": rem(8),
                "--input-fz": rem(18),
              },
            };
          default:
            return { wrapper: {} };
        }
      },
    }),
    FileInput: FileInput.extend({
      vars: (_, props) => {
        switch (props.size) {
          case "md":
            return {
              wrapper: {
                "--input-height": rem(42),
                "--input-padding-x": rem(8),
                "--input-fz": rem(18),
              },
            };
          default:
            return { wrapper: {} };
        }
      },
    }),

    Button: Button.extend({
      classNames: {
        root: classes.buttonRoot,
      },
    }),

    Input: Input.extend({
      classNames: {
        wrapper: classes.inputWrapper,
        input: classes.input,
      },
    }),

    PasswordInput: Input.extend({
      classNames: {
        input: classes.passInput,
      },
      vars: (_, props) => {
        switch (props.size) {
          case "md":
            return {
              wrapper: {
                "--input-height": rem(42),
                "--input-padding-x": rem(8),
                "--input-fz": rem(18),
              },
            };
          default:
            return { wrapper: {} };
        }
      },
    }),

    InputWrapper: Input.Wrapper.extend({
      classNames: {
        root: classes.inputWrapperRoot,
        label: classes.label,
        error: classes.error,
      },
    }),

    Slider: RangeSlider.extend({
      classNames: {
        label: classes.thumbLabel,
        markLabel: classes.markLabel,
      },
    }),
  },
  other: {
    fontWeights: {
      regular: 400,
      medium: 500,
      bold: 700,
    },
  },
});
