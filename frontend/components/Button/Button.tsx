import { ButtonProps } from "@mantine/core";

type Props = ButtonProps;
export const Button = (props: Props) => {
  return <Button {...props}>{props.children}</Button>;
};
