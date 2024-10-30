import { Group, LoadingOverlay, ScrollArea, Stack, Title } from "@mantine/core";
import { Recipe, Props as RecipeProps } from "../Recipe";
import { memo } from "react";

type Props = {
  items?: RecipeProps[];
  isFetching?: boolean;
};

export const RecipeList = memo(
  ({
    items = Array.from({ length: 10 }, () => ({
      name: "Рецепт",
      time: "1 20",
    })),
    isFetching = true,
  }: Props) => (
    <Stack gap={20}>
      <Title c="accent.0" order={1}>
        Рецепты
      </Title>

      <LoadingOverlay
        visible={isFetching}
        zIndex={1000}
        overlayProps={{ radius: "sm", blur: 2 }}
        loaderProps={{ color: "primary.0", type: "bars" }}
      />

      <ScrollArea type="never" w="100%">
        <Group wrap="nowrap" gap={20}>
          {items.map((item) => (
            <Recipe key={item.name} {...item} />
          ))}
        </Group>
      </ScrollArea>
    </Stack>
  )
);
