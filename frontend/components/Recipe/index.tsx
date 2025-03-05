import { Text, Title, Stack, BackgroundImage, Box } from "@mantine/core";

import styles from "./Recipe.module.css";
import { Trash } from "../icons/Trash";
import { ItemCardIcon } from "../ItemCard/ItemCardIcon";
import { showErrorNotification } from "@/utils";
import { AxiosError } from "axios";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { deleteRecipe } from "@/api/recipes/deleteRecipe";
import { getRecipes } from "@/api/recipes/getRecipes";

export type Props = {
  uid?: string;
  name: string;
  time: string;
  onClick?: () => void;
  img?: string;
  editable?: boolean;
  onExtended?: () => void;
};

export const Recipe = ({
  uid,
  onClick,
  name,
  time,
  img,
  editable,
  onExtended,
}: Props) => {
  const fallbackImg =
    "https://raw.githubusercontent.com/mantinedev/mantine/master/.demo/images/bg-8.png";

  const queryClient = useQueryClient();

  const { mutate: mutateDelete } = useMutation({
    mutationFn: deleteRecipe,
    mutationKey: [deleteRecipe.queryKey],
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: [getRecipes.queryKey],
      });
    },
    onError: (error: AxiosError<any>) => {
      showErrorNotification(error);
    },
  });

  return (
    <BackgroundImage
      mih={280}
      w={200}
      display="flex"
      src={img ?? fallbackImg}
      radius="lg"
      pos="relative"
      className={styles.root}
      onClick={onClick}
    >
      <Stack w="100%" className={styles.main} gap={16} justify="flex-end">
        <Box py={20} px={16}>
          <Title order={3} c="accent.0">
            {name}
          </Title>

          {editable && uid && (
            <Box
              className={styles.deleteIcon}
              onClick={() => mutateDelete(uid)}
            >
              <Trash />
            </Box>
          )}

          {onExtended && (
            <ItemCardIcon type="max" onClick={() => onExtended?.()} />
          )}

          <Text fw={500} fz={18} c="accent.1">
            {time}
          </Text>
        </Box>
      </Stack>
    </BackgroundImage>
  );
};
