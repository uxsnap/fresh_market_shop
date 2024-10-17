import { Glass } from "@/components/icons/Glass";
import { StyleProp, TextInput } from "@mantine/core";

type Props = {
  className?: string;
  maw?: StyleProp<React.CSSProperties['maxWidth']>;
};

export const Search = ({ className, maw = 400 }: Props) => {
  return (
    <TextInput
      className={className}
      w="100%"
      maw={maw}
      size="md"
      leftSection={<Glass size={16} />}
      placeholder="Поиск товаров"
    />
  );
};
