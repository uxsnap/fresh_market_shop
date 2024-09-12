import { IconProps } from "@/types";
import { rem } from "@mantine/core";

export const ArrowDown = ({ size, fill, ...other }: IconProps) => {
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
      <path
        d="M12 13.7858L17.9531 7.82803C18.3937 7.3874 19.1062 7.3874 19.5422 7.82803C19.9781 8.26865 19.9781 8.98115 19.5422 9.42178L12.7969 16.1718C12.3703 16.5983 11.6859 16.6077 11.2453 16.2046L4.45312 9.42647C4.23281 9.20615 4.125 8.91553 4.125 8.62959C4.125 8.34365 4.23281 8.05303 4.45312 7.83272C4.89375 7.39209 5.60625 7.39209 6.04219 7.83272L12 13.7858Z"
        fill="black"
      />
    </svg>
  );
};
