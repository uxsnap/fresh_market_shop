import { useAdminStore } from "@/store/admin";
import { Recipe as IRecipe } from "@/types";
import { formatDuration, getRecipeBg } from "@/utils";
import { Group, LoadingOverlay, Pagination, Stack } from "@mantine/core";
import { useIsFirstRender, useWindowScroll } from "@mantine/hooks";
import { useQuery } from "@tanstack/react-query";
import { useEffect, useState } from "react";
import { Filters } from "./Filters";
import { getRecipes } from "@/api/recipes/getRecipes";
import { Recipe } from "@/components/Recipe";

const RECIPES_LIST_LIMIT = 30;

export const AdminRecipeList = () => {
  const [activePage, setPage] = useState(1);
  const [_, scrollTo] = useWindowScroll();
  const firstRender = useIsFirstRender();

  const [filters, setFilters] = useState({ name: "" });

  const setModalOpen = useAdminStore((s) => s.setModalOpen);
  const setRecipeItem = useAdminStore((s) => s.setRecipeItem);

  const { data, isLoading } = useQuery({
    queryKey: [getRecipes.queryKey, activePage, filters],
    queryFn: () =>
      getRecipes({
        page: activePage,
        limit: RECIPES_LIST_LIMIT,
        ...filters,
      }),
  });

  useEffect(() => {
    scrollTo({ y: 0 });
  }, [activePage]);

  const handleEdit = (item: IRecipe) => {
    setRecipeItem(item);
    setModalOpen(true);
  };

  return (
    <Stack gap={20} pos="relative">
      <LoadingOverlay
        visible={isLoading && firstRender}
        zIndex={1}
        overlayProps={{ radius: "sm", blur: 2 }}
        loaderProps={{ color: "primary.0", type: "bars" }}
      />

      <Filters setFilters={setFilters} />

      <Group mt={12} gap={12} justify="left">
        {data?.data.recipes.map((item) => (
          <Recipe
            key={item.uid}
            name={item.name}
            img={getRecipeBg(item.uid)}
            time={formatDuration(item.cookingTime)}
          />
        ))}
      </Group>

      {data?.data.total && (
        <Pagination
          color="accent.0"
          value={activePage}
          onChange={setPage}
          total={Math.ceil(data.data.total / RECIPES_LIST_LIMIT)}
        />
      )}
    </Stack>
  );
};
