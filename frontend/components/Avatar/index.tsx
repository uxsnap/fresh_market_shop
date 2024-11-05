import { Avatar as MantineAvatar } from "@mantine/core";

import { User } from "../icons/User";

type Props = {
  src?: string;
  size?: "small" | "default";
  onClick?: () => void;
};

export const Avatar = ({ src = "", size = "default", onClick }: Props) => {
  return (
    <MantineAvatar
      onClick={onClick}
      src={src}
      c="bg.1"
      size={size === "default" ? 250 : 38}
    >
      <User
        fill={`var(--mantine-color-accent-2`}
        size={size === "default" ? 100 : 19}
      />
    </MantineAvatar>
  );
};
