import { IconProps } from "@/types";
import { rem } from "@mantine/core";

export const MailSent = ({ size, fill, ...other }: IconProps) => {
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
        d="M3 7H6"
        stroke={fill}
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
      />
      <path
        d="M3 11H5"
        stroke={fill}
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
      />
      <path
        d="M9.01996 8.801L8.41996 14.801C8.39216 15.079 8.42292 15.3598 8.51027 15.6252C8.59762 15.8906 8.73962 16.1348 8.92711 16.342C9.11459 16.5492 9.34342 16.7148 9.59882 16.8281C9.85422 16.9414 10.1305 17 10.41 17H18.39C18.8859 17 19.3642 16.8157 19.7319 16.4829C20.0997 16.1501 20.3306 15.6925 20.38 15.199L20.98 9.199C21.0078 8.92097 20.977 8.64019 20.8896 8.37478C20.8023 8.10936 20.6603 7.86519 20.4728 7.65801C20.2853 7.45083 20.0565 7.28524 19.8011 7.17191C19.5457 7.05857 19.2694 7.00001 18.99 7H11.01C10.514 7.00002 10.0357 7.18432 9.66799 7.51712C9.30026 7.84993 9.06931 8.30749 9.01996 8.801Z"
        stroke={fill}
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
      />
      <path
        d="M9.80005 7.5L12.782 10.78C13.047 11.0715 13.3669 11.308 13.7233 11.476C14.0797 11.6439 14.4657 11.74 14.8592 11.7588C15.2528 11.7775 15.6461 11.7186 16.0169 11.5853C16.3876 11.452 16.7285 11.247 17.02 10.982L20.3 8"
        stroke={fill}
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
      />
    </svg>
  );
};
