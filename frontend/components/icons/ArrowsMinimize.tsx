import { IconProps } from "@/types";
import { rem } from "@mantine/core";

export const ArrowsMinimize = ({ size, fill, ...other }: IconProps) => {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill={fill}
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
      style={{ width: rem(size), height: rem(size) }}
      {...other}
    >
      <path stroke="none" d="M0 0h24v24H0z" fill="none" />
      <path d="M5 9l4 0l0 -4" />
      <path d="M3 3l6 6" />
      <path d="M5 15l4 0l0 4" />
      <path d="M3 21l6 -6" />
      <path d="M19 9l-4 0l0 -4" />
      <path d="M15 9l6 -6" />
      <path d="M19 15l-4 0l0 4" />
      <path d="M15 15l6 6" />
    </svg>
  );
};
