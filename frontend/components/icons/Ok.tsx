import { IconProps } from "@/types";
import { rem } from "@mantine/core";

export const Ok = ({ size, fill, ...other }: IconProps) => {
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
      <g clipPath="url(#clip0_72_5643)">
        <path d="M8.15158 20.5971L0.351579 12.7971C-0.11703 12.3285 -0.11703 11.5687 0.351579 11.1L2.04859 9.40297C2.5172 8.93431 3.27705 8.93431 3.74566 9.40297L9.00011 14.6574L20.2546 3.40297C20.7232 2.93436 21.483 2.93436 21.9516 3.40297L23.6486 5.10003C24.1173 5.56864 24.1173 6.32844 23.6486 6.7971L9.84864 20.5971C9.37999 21.0658 8.62019 21.0658 8.15158 20.5971Z" />
      </g>
      <defs>
        <clipPath id="clip0_72_5643">
          <rect width="24" height="24" fill="white" />
        </clipPath>
      </defs>
    </svg>
  );
};
