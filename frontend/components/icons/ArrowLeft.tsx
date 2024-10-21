import { IconProps } from "@/types";
import { rem } from "@mantine/core";

export const ArrowLeft = ({ size, fill, ...other }: IconProps) => {
  return (
    <svg
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      style={{ width: rem(size), height: rem(size) }}
      {...other}
    >
      <path
        d="M11.4375 18.75L4.6875 12L11.4375 5.25M5.625 12H19.3125"
        stroke={fill}
        strokeWidth="2.25"
        strokeLinecap="round"
        strokeLinejoin="round"
      />
    </svg>
  );
};
