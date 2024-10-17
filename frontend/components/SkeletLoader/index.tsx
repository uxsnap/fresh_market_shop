import { Skeleton } from "@mantine/core";

type Props = {
  h?: number;
  l?: number;
};

export const SkeletLoader = ({ h = 32, l = 5 }: Props) => {
  return Array.from({ length: l }, (_, ind) => <Skeleton height={h + "px"} key={ind} radius="md" />);
};
