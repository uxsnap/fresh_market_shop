import { editProduct } from "@/api/products/editProduct";
import { getProducts } from "@/api/products/getProducts";
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
  Input,
  NumberInput,
  Select,
  Stack,
  Textarea,
  TextInput,
} from "@mantine/core";
import { hasLength, isNotEmpty, useForm } from "@mantine/form";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { useEffect, useState } from "react";
import { ImgsUpload } from "../ImgsUpload";
import { FileWithPath } from "@mantine/dropzone";
import { updatePhotos } from "@/api/products/updatePhotos";
import { BackendImg } from "@/types";
import { createRecipe } from "@/api/recipes/createRecipe";
import { TimeInput } from "@mantine/dates";

import styles from "./CreateButton.module.css";
import { getRecipes } from "@/api/recipes/getRecipes";
import { IMaskInput } from "react-imask";
import { COOKING_TIME_BORDERS } from "@/constants";

type Props = {
  onClose: () => void;
};

const weightData = [
  { label: "20 грамм", value: "20" },
  { label: "100 грамм", value: "100" },
];

export const RecipeModal = ({ onClose }: Props) => {
  const recipeItem = useAdminStore((s) => s.recipeItem);
  const setRecipeItem = useAdminStore((s) => s.setRecipeItem);

  // const [files, setFiles] = useState<(FileWithPath | BackendImg)[]>([]);

  // useEffect(() => {
  //   if (!productItem?.imgs.length) {
  //     return;
  //   }

  //   setFiles(productItem.imgs);
  // }, [productItem]);

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
    mutationFn: editProduct,
    onSuccess: () => {
      onClose();

      queryClient.invalidateQueries({
        queryKey: [getRecipes.queryKey],
      });

      showSuccessNotification("Продукт успешно обновлен!");
    },
    onError: (error: AxiosError<any>) => {
      showErrorNotification(error);
    },
  });

  const { mutate: mutateFiles } = useMutation({
    mutationFn: updatePhotos,
    onError: (error: AxiosError<any>) => {
      showErrorNotification(error);
    },
  });

  // const handleFiles = () => {
  //   if (!files.length || !productItem) {
  //     return;
  //   }

  //   const form = new FormData();
  //   form.append("category", productItem.categoryUid);
  //   form.append("uid", productItem.id);

  //   for (const file of files) {
  //     if ("uid" in file) {
  //       continue;
  //     }

  //     form.append("file", file);
  //   }

  //   mutateFiles(form);
  // };

  const handleSubmit = form.onSubmit((values) => {
    const submitValues = {
      ...values,
      cookingTime: convertTimeToDuration(values.cookingTime),
    };

    // if (recipeItem) {
    //   mutateEdit({ ...submitValues, uid: recipeItem.id });
    // } else {
    mutateCreate(submitValues);
    // }

    // handleFiles();
  });

  return (
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
          key={form.key("cookingTime")}
          {...form.getInputProps("cookingTime")}
          classNames={{ input: styles.time }}
        />

        {/* <Input.Error {...form.getInputProps("cookingTime")} /> */}
        {/* </Input.Wrapper> */}

        {/* {productItem && (
          <ImgsUpload
            productUid={productItem.id}
            files={files}
            setFiles={setFiles}
          />
        )} */}

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
  );
};
