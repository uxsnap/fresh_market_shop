import { IconProps } from "@/types";
import { rem } from "@mantine/core";

export const Refresh = ({ size = 24, fill, ...other }: IconProps) => {
  return (
    <svg
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill={fill}
      xmlns="http://www.w3.org/2000/svg"
      style={{ width: rem(size), height: rem(size) }}
      {...other}
    >
      <g clipPath="url(#clip0_49_1298)">
        <path
          d="M17.65 6.35C16.9099 5.60485 16.0296 5.01356 15.0599 4.61023C14.0902 4.2069 13.0503 3.99951 12 4C7.58001 4 4.01001 7.58 4.01001 12C4.01001 16.42 7.58001 20 12 20C15.73 20 18.84 17.45 19.73 14H17.65C17.2381 15.1695 16.4734 16.1824 15.4614 16.8988C14.4494 17.6153 13.24 18 12 18C8.69001 18 6.00001 15.31 6.00001 12C6.00001 8.69 8.69001 6 12 6C13.66 6 15.14 6.69 16.22 7.78L13 11H20V4L17.65 6.35Z"
          fill="black"
        />
      </g>
      <defs>
        <clipPath id="clip0_49_1298">
          <rect width="24" height="24" fill="white" />
        </clipPath>
      </defs>
    </svg>
  );
};
