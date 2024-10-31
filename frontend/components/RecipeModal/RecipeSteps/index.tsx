import { useQuery } from "@tanstack/react-query";
import { RecipeStep } from "../RecipeStep";
import { getRecipeSteps } from "@/api/recipes/getRecipeSteps";
import { Box, LoadingOverlay, Stack } from "@mantine/core";
import { getRecipeStepImg } from "@/utils";

type Props = {
  uid: string;
};

export const RecipeSteps = ({ uid }: Props) => {
  const { data, isFetching } = useQuery({
    queryFn: () => getRecipeSteps(uid),
    queryKey: [getRecipeSteps.queryKey, uid],
  });

  return (
    <Stack gap={12} pos="relative">
      <LoadingOverlay
        visible={isFetching}
        zIndex={1}
        overlayProps={{ radius: "sm", blur: 2 }}
        loaderProps={{ color: "primary.0", type: "bars" }}
      />

      {data?.data.map((step) => (
        <RecipeStep
          src={getRecipeStepImg(step)}
          key={step.recipeUid}
          step={step.step}
          maxStep={data.data.length}
        >
          {step.description}
        </RecipeStep>
      ))}
    </Stack>
  );
};
