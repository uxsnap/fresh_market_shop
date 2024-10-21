import { IconProps } from "@/types";
import { rem } from "@mantine/core";

export const Trash = ({ size, fill, ...other }: IconProps) => {
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
      <g clipPath="url(#clip0_58_1320)">
        <path
          d="M21.75 1.50001H16.125L15.6844 0.623446C15.591 0.436046 15.4473 0.27841 15.2692 0.16827C15.0912 0.0581305 14.8859 -0.000141936 14.6766 8.21846e-06H9.31875C9.10987 -0.000794775 8.90498 0.0572604 8.72756 0.167522C8.55015 0.277784 8.40739 0.435793 8.31562 0.623446L7.875 1.50001H2.25C2.05109 1.50001 1.86032 1.57903 1.71967 1.71968C1.57902 1.86033 1.5 2.0511 1.5 2.25001L1.5 3.75001C1.5 3.94892 1.57902 4.13969 1.71967 4.28034C1.86032 4.42099 2.05109 4.50001 2.25 4.50001H21.75C21.9489 4.50001 22.1397 4.42099 22.2803 4.28034C22.421 4.13969 22.5 3.94892 22.5 3.75001V2.25001C22.5 2.0511 22.421 1.86033 22.2803 1.71968C22.1397 1.57903 21.9489 1.50001 21.75 1.50001ZM3.99375 21.8906C4.02952 22.4619 4.28164 22.998 4.69878 23.3899C5.11591 23.7817 5.66672 23.9999 6.23906 24H17.7609C18.3333 23.9999 18.8841 23.7817 19.3012 23.3899C19.7184 22.998 19.9705 22.4619 20.0062 21.8906L21 6.00001H3L3.99375 21.8906Z"
        />
      </g>
      <defs>
        <clipPath id="clip0_58_1320">
          <rect width="24" height="24" fill="white" />
        </clipPath>
      </defs>
    </svg>
  );
};
