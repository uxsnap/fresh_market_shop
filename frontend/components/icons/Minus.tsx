import { IconProps } from "@/types";
import { rem } from "@mantine/core";

export const Minus = ({ size, fill, ...others }: IconProps) => {
  return (
    <svg
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill={fill}
      xmlns="http://www.w3.org/2000/svg"
      style={{ width: rem(size), height: rem(size) }}
      {...others}
    >
      <path d="M21 9.75H3C2.17172 9.75 1.5 10.4217 1.5 11.25V12.75C1.5 13.5783 2.17172 14.25 3 14.25H21C21.8283 14.25 22.5 13.5783 22.5 12.75V11.25C22.5 10.4217 21.8283 9.75 21 9.75Z" />
    </svg>
  );
};
