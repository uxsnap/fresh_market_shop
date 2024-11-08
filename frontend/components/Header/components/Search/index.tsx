import { search } from "@/api/search";
import { CategoryItem } from "@/components/CategoryItem";
import { Glass } from "@/components/icons/Glass";
import { PopOver } from "@/components/PopOver";
import { SkeletLoader } from "@/components/SkeletLoader";
import { SmallCartItem } from "@/components/SmallCartItem";
import { StyleProp, TextInput } from "@mantine/core";
import { useDebouncedValue } from "@mantine/hooks";
import { useQuery } from "@tanstack/react-query";
import { useState } from "react";
import cn from "classnames";
import { Text } from "@mantine/core";
import { useRouter } from "next/navigation";

import styles from "./Search.module.css";

type Props = {
  className?: string;
  maw?: StyleProp<React.CSSProperties["maxWidth"]>;
};

export const Search = ({ className, maw = 400 }: Props) => {
  const [name, setName] = useState("");
  const [debounced] = useDebouncedValue(name, 200);

  const router = useRouter();

  const { data, isFetching, isFetched } = useQuery({
    queryKey: [search.queryKey, debounced],
    queryFn: () => search(debounced),
    enabled: !!debounced,
    retry: 0,
  });

  const renderLoader = () => <SkeletLoader h={40} l={4} />;

  const renderData = () => {
    const { products = [], categories = [] } = data?.data ?? {};

    if (!products.length && !categories.length) {
      return (
        <Text c="accent.0" fz={18} style={{ textAlign: "center" }} fw={600}>
          Нет результатов
        </Text>
      );
    }

    return [
      ...categories?.map((category) => (
        <CategoryItem
          key={category.uid}
          onClick={() => router.push(`/products/${category.uid}`)}
        >
          {category.name}
        </CategoryItem>
      )),
      ...products.map((product) => (
        <SmallCartItem key={product.uid} img={product.img}>
          {product.name}
        </SmallCartItem>
      )),
    ];
  };

  return (
    <div className={cn(styles.root, className)}>
      <TextInput
        pos="relative"
        w="100%"
        maw={maw}
        size="md"
        className={styles.input}
        leftSection={<Glass size={16} />}
        placeholder="Поиск товаров и категорий"
        value={name}
        onChange={(event) => setName(event.currentTarget.value)}
      />

      {name && (
        <PopOver className={cn(styles.popover, !data?.data && styles.loading)}>
          {isFetching && renderLoader()}

          {isFetched && renderData()}
        </PopOver>
      )}
    </div>
  );
};
