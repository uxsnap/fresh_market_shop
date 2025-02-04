import { useAdminStore } from "@/store/admin";
import {
  convertDurationToTime,
  convertTimeToDuration,
  showErrorNotification,
  showSuccessNotification,
} from "@/utils";
import {
  Button,
  Group,
  Modal,
  NumberInput,
  Stack,
  TextInput,
  Title,
} from "@mantine/core";
import { hasLength, isNotEmpty, useForm } from "@mantine/form";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { useEffect, useState } from "react";
import { createRecipe } from "@/api/recipes/createRecipe";

import styles from "./RecipeModal.module.css";
import { getRecipes } from "@/api/recipes/getRecipes";
import { IMaskInput } from "react-imask";
import { COOKING_TIME_BORDERS } from "@/constants";
import { editRecipe } from "@/api/recipes/editRecipe";
import { StepsModal } from "./StepsModal";
import { RecipeStepObj } from "@/types";

type Props = {
  onClose: () => void;
};

export const RecipeModal = ({ onClose }: Props) => {
  const recipeItem = useAdminStore((s) => s.recipeItem);
  const setRecipeItem = useAdminStore((s) => s.setRecipeItem);

  const [stepsModalOpened, setStepsModalOpened] = useState(false);
  const [curSteps, setCurSteps] = useState<RecipeStepObj[]>([]);

  const queryClient = useQueryClient();

  const form = useForm({
    mode: "uncontrolled",
    initialValues: {
      name: "",
      ccal: 0,
      cookingTime: "",
    },
    validate: {
      name: hasLength({ min: 1 }, "Длина названия должна быть больше 1"),
      ccal: isNotEmpty("Калорийность не должна быть пустой"),
      cookingTime: (value) => {
        if (value.length < 5) {
          return "Время приготовления не должно быть пустым";
        }

        const time = convertTimeToDuration(value);

        if (
          time < COOKING_TIME_BORDERS.min ||
          time > COOKING_TIME_BORDERS.max
        ) {
          return "Время приготовления выходит за временные границы";
        }

        return null;
      },
    },
  });

  useEffect(() => {
    if (!recipeItem) {
      return;
    }

    form.reset();

    form.setValues({
      name: recipeItem.name,
      cookingTime: convertDurationToTime(recipeItem.cookingTime),
      ccal: recipeItem.ccal,
    });
  }, [recipeItem]);

  const handleClose = () => {
    setRecipeItem(undefined);
    onClose();
  };

  const { mutate: mutateCreate, isPending: isPendingCreate } = useMutation({
    mutationFn: createRecipe,
    onSuccess: () => {
      onClose();

      queryClient.invalidateQueries({
        queryKey: [getRecipes.queryKey],
      });

      showSuccessNotification("Рецепт успешно добавлен!");
    },
    onError: (error: AxiosError<any>) => {
      showErrorNotification(error);
    },
  });

  const { mutate: mutateEdit, isPending: isPendingUpdate } = useMutation({
    mutationFn: editRecipe,
    onSuccess: () => {
      onClose();

      queryClient.invalidateQueries({
        queryKey: [getRecipes.queryKey],
      });

      showSuccessNotification("Рецепт успешно обновлен!");
    },
    onError: (error: AxiosError<any>) => {
      showErrorNotification(error);
    },
  });

  const handleSubmit = form.onSubmit((values) => {
    const submitValues = {
      ...values,
      cookingTime: convertTimeToDuration(values.cookingTime),
    };

    if (recipeItem) {
      mutateEdit({ ...submitValues, uid: recipeItem.uid });
    } else {
      mutateCreate(submitValues);
    }
  });

  return (
    <>
      <Modal
        opened={stepsModalOpened}
        onClose={() => setStepsModalOpened(false)}
        title={
          <Title c="accent.0" order={4}>
            Шаги рецепта
          </Title>
        }
      >
        <StepsModal
          onClose={() => setStepsModalOpened(false)}
          steps={curSteps}
          onChange={setCurSteps}
        />
      </Modal>

      <form onSubmit={handleSubmit}>
        <Stack gap={16}>
          <TextInput
            size="md"
            label="Название"
            withAsterisk
            required
            placeholder="Введите название"
            {...form.getInputProps("name")}
          />

          <NumberInput
            w="100%"
            min={1}
            hideControls
            allowLeadingZeros={false}
            allowNegative={false}
            allowDecimal={false}
            withAsterisk
            lh={1}
            size="md"
            required
            label="Калории"
            placeholder="Введите калории"
            key={form.key("ccal")}
            {...form.getInputProps("ccal")}
          />

          <TextInput
            placeholder="Введите время приготовления"
            component={IMaskInput}
            // @ts-ignore
            mask="00:00"
            w="100%"
            lh={1}
            required
            size="md"
            withAsterisk
            label="Время приготовления"
            classNames={{ input: styles.time }}
            key={form.key("cookingTime")}
            {...form.getInputProps("cookingTime")}
          />

          <Button
            disabled={isPendingCreate || isPendingUpdate}
            w="100%"
            variant="dashed"
            onClick={() => setStepsModalOpened(true)}
          >
            Редактировать шаги
          </Button>

          <Group wrap="nowrap" mt={4} justify="space-between">
            <Button
              disabled={isPendingCreate || isPendingUpdate}
              w="100%"
              type="submit"
              variant="accent"
            >
              Сохранить
            </Button>

            <Button w="100%" onClick={handleClose} variant="accent-reverse">
              Закрыть
            </Button>
          </Group>
        </Stack>
      </form>
    </>
  );
};
