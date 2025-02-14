import { useAdminStore } from "@/store/admin";
import {
  convertDurationToTime,
  convertTimeToDuration,
  getRecipeBg,
  processImgFile,
  renameFile,
  showErrorNotification,
  showSuccessNotification,
  urlToObject,
} from "@/utils";
import {
  Button,
  FileInput,
  Group,
  Modal,
  NumberInput,
  Stack,
  TextInput,
  Title,
  Image,
} from "@mantine/core";
import { hasLength, isNotEmpty, useForm } from "@mantine/form";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { AxiosError, AxiosResponse } from "axios";
import { useEffect, useState } from "react";
import { createRecipe } from "@/api/recipes/createRecipe";

import styles from "./RecipeModal.module.css";
import { getRecipes } from "@/api/recipes/getRecipes";
import { IMaskInput } from "react-imask";
import { editRecipe } from "@/api/recipes/editRecipe";
import { StepsModal } from "./StepsModal";
import { Recipe, RecipeStepObj } from "@/types";
import { deleteRecipePhotos } from "@/api/recipes/deleteRecipePhotos";
import { addRecipePhotos } from "@/api/recipes/addRecipePhotos";
import { validateCookingTime } from "./utils";

type Props = {
  onClose: () => void;
};

type Form = {
  name: string;
  ccal: number;
  cookingTime: string;
  img: null | File;
};

export const RecipeModal = ({ onClose }: Props) => {
  const recipeItem = useAdminStore((s) => s.recipeItem);
  const setRecipeItem = useAdminStore((s) => s.setRecipeItem);

  const [stepsModalOpened, setStepsModalOpened] = useState(false);

  const queryClient = useQueryClient();

  const form = useForm<Form>({
    mode: "uncontrolled",
    initialValues: {
      name: "",
      ccal: 0,
      cookingTime: "",
      img: null,
    },
    validate: {
      name: hasLength({ min: 1 }, "Длина названия должна быть больше 1"),
      ccal: isNotEmpty("Калорийность не должна быть пустой"),
      cookingTime: validateCookingTime,
    },
  });

  useEffect(() => {
    const asyncFunc = async () => {
      form.reset();

      if (!recipeItem) {
        return;
      }

      const img = await urlToObject(getRecipeBg(recipeItem.uid), "0.webp");

      form.setValues({
        name: recipeItem.name,
        cookingTime: convertDurationToTime(recipeItem.cookingTime),
        ccal: recipeItem.ccal,
        img,
      });
    };

    asyncFunc();
  }, []);

  const handleClose = () => {
    setRecipeItem(undefined);

    onClose();
  };

  const { mutate: mutateCreate, isPending: isPendingCreate } = useMutation({
    mutationFn: createRecipe,
    onSuccess: ({ data }: AxiosResponse<Recipe>) => {
      showSuccessNotification("Рецепт успешно добавлен!");

      setRecipeItem(data);

      handleMainPhoto(data.uid, form.getValues().img);
    },
    onError: (error: AxiosError<any>) => {
      showErrorNotification(error);
    },
  });

  const { mutate: mutateEdit, isPending: isPendingUpdate } = useMutation({
    mutationFn: editRecipe,
    onSuccess: ({ data }: AxiosResponse<{ uid: string }>) => {
      showSuccessNotification("Рецепт успешно обновлен!");

      queryClient.invalidateQueries({
        queryKey: [getRecipes.queryKey],
      });

      handleMainPhoto(data.uid, form.getValues().img);
    },
    onError: (error: AxiosError<any>) => {
      showErrorNotification(error);
    },
  });

  const { mutate: mutateDeletePhoto, isPending: isPendingDeletePhoto } =
    useMutation({
      mutationFn: deleteRecipePhotos,
      onError: (error: AxiosError<any>) => {
        showErrorNotification(error);
      },
    });

  const { mutate: mutateAddPhoto, isPending: isPendingAddPhoto } = useMutation({
    mutationFn: addRecipePhotos,
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: [getRecipes.queryKey],
      });
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

  const handleMainPhoto = async (uid: string, img: File | null) => {
    if (!img) {
      mutateDeletePhoto({ uid, photos: ["0.webp"] });
      return;
    }

    const processed = await processImgFile(img);

    const formData = new FormData();
    formData.set("uid", uid);
    formData.set("file", renameFile(processed, "0.webp"));

    mutateAddPhoto(formData);
  };

  const isPending =
    isPendingCreate ||
    isPendingUpdate ||
    isPendingDeletePhoto ||
    isPendingAddPhoto;

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
        {recipeItem && (
          <StepsModal
            uid={recipeItem.uid}
            onClose={() => setStepsModalOpened(false)}
          />
        )}
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
            label="Время приготовления (ЧЧ:ММ)"
            classNames={{ input: styles.time }}
            key={form.key("cookingTime")}
            {...form.getInputProps("cookingTime")}
          />

          <FileInput
            size="md"
            label="Изображение"
            withAsterisk
            accept="image/webp"
            required
            placeholder="Выберите основное изображение"
            clearable
            key={form.key(`img`)}
            {...form.getInputProps(`img`)}
            onChange={(payload) => form.setFieldValue(`img`, payload)}
          />

          {form.getValues().img && (
            <Image
              style={{ width: 100, height: 100 }}
              src={URL.createObjectURL(form.getValues().img as File)}
            />
          )}

          <Button
            disabled={isPending || !recipeItem}
            w="100%"
            variant="dashed"
            onClick={() => setStepsModalOpened(true)}
          >
            Редактировать шаги
          </Button>

          <Group wrap="nowrap" mt={4} justify="space-between">
            <Button
              disabled={isPending}
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
