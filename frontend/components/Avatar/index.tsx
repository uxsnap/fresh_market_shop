import { Avatar as MantineAvatar } from "@mantine/core";

import { User } from "../icons/User";

type Props = {
  src?: string;
};

export const Avatar = ({ src = "" }: Props) => {
  return (
    <MantineAvatar src={src} c="bg.1" size={250}>
      <User fill={`var(--mantine-color-accent-2`} size={100} />
    </MantineAvatar>
  );
};
