import { IconProps } from "@/types";
import { rem } from "@mantine/core";

export const Hourglass = ({ size, fill, ...other }: IconProps) => {
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
      <g clipPath="url(#clip0_553_2509)">
        <path
          d="M6 20V18C6 16.4087 6.63214 14.8826 7.75736 13.7574C8.88258 12.6321 10.4087 12 12 12C13.5913 12 15.1174 12.6321 16.2426 13.7574C17.3679 14.8826 18 16.4087 18 18V20C18 20.2652 17.8946 20.5196 17.7071 20.7071C17.5196 20.8946 17.2652 21 17 21H7C6.73478 21 6.48043 20.8946 6.29289 20.7071C6.10536 20.5196 6 20.2652 6 20Z"
          stroke={fill}
          strokeWidth="2"
          strokeLinecap="round"
          strokeLinejoin="round"
        />
        <path
          d="M6 4V6C6 7.5913 6.63214 9.11742 7.75736 10.2426C8.88258 11.3679 10.4087 12 12 12C13.5913 12 15.1174 11.3679 16.2426 10.2426C17.3679 9.11742 18 7.5913 18 6V4C18 3.73478 17.8946 3.48043 17.7071 3.29289C17.5196 3.10536 17.2652 3 17 3H7C6.73478 3 6.48043 3.10536 6.29289 3.29289C6.10536 3.48043 6 3.73478 6 4Z"
          stroke={fill}
          strokeWidth="2"
          strokeLinecap="round"
          strokeLinejoin="round"
        />
      </g>
      <defs>
        <clipPath id="clip0_553_2509">
          <rect width="24" height="24" fill="white" />
        </clipPath>
      </defs>
    </svg>
  );
};
