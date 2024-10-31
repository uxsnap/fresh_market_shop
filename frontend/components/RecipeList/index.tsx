import { Group, LoadingOverlay, ScrollArea, Stack, Title } from "@mantine/core";
import { Recipe } from "../Recipe";
import { memo, useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { getRecipes } from "@/api/recipes/getRecipes";

import { formatDuration, getRecipeBg } from "@/utils";
import { RecipeModal } from "../RecipeModal";
import { Recipe as IRecipe } from "@/types";

export const RecipeList = memo(() => {
  const { data, isFetching } = useQuery({
    queryFn: getRecipes,
    queryKey: [getRecipes.queryKey],
  });

  const [curRecipe, setCurRecipe] = useState<IRecipe>();

  return (
    <>
      {curRecipe && (
        <RecipeModal
          close={() => setCurRecipe(undefined)}
          uid={curRecipe.uid}
          name={curRecipe.name}
          ccal={curRecipe.ccal}
          img={getRecipeBg(curRecipe.uid)}
        />
      )}

      <Stack mih={280} gap={20} pos="relative">
        <Title c="accent.0" order={1}>
          Рецепты
        </Title>

        <LoadingOverlay
          visible={isFetching}
          zIndex={1}
          overlayProps={{ radius: "sm", blur: 2 }}
          loaderProps={{ color: "primary.0", type: "bars" }}
        />

        <ScrollArea type="never" w="100%">
          <Group wrap="nowrap" gap={20}>
            {data?.data.map((item) => (
              <Recipe
                key={item.uid}
                name={item.name}
                onClick={() => setCurRecipe(item)}
                img={getRecipeBg(item.uid)}
                time={formatDuration(item.cookingTime)}
              />
            ))}
          </Group>
        </ScrollArea>
      </Stack>
    </>
  );
});
