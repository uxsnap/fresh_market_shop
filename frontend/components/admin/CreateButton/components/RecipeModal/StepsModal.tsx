import { addRecipePhotos } from "@/api/recipes/addRecipePhotos";
import { addRecipeSteps } from "@/api/recipes/addRecipeSteps";
import { deleteRecipePhotos } from "@/api/recipes/deleteRecipePhotos";
import { getRecipeSteps } from "@/api/recipes/getRecipeSteps";
import { RecipeStepObj } from "@/types";
import {
  getRecipeStepImg,
  processImgFile,
  renameFile,
  showErrorNotification,
  showSuccessNotification,
  urlToObject,
} from "@/utils";
import {
  ActionIcon,
  Button,
  FileInput,
  Group,
  Paper,
  Stack,
  Textarea,
  Image,
} from "@mantine/core";
import { isNotEmpty, useForm } from "@mantine/form";
import { IconTrash } from "@tabler/icons-react";
import { useMutation, useQuery } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { useEffect } from "react";

export type Props = {
  uid: string;
  onClose: () => void;
};

export const StepsModal = ({ uid, onClose }: Props) => {
  const form = useForm({
    mode: "uncontrolled",
    initialValues: {
      formSteps: [] as RecipeStepObj[],
    },
    validate: {
      formSteps: {
        description: isNotEmpty("Длина описания должна быть больше 0"),
        img: isNotEmpty("У шага должна быть фотография"),
      },
    },
  });

  const { data } = useQuery({
    queryFn: () => getRecipeSteps(uid!),
    enabled: !!uid,
    queryKey: [getRecipeSteps.queryKey, uid],
  });

  const { mutate: mutateAdd, isPending: isPendingAdding } = useMutation({
    mutationFn: addRecipeSteps,
    onSuccess: () => {
      showSuccessNotification("Шаги приготовления успешно обновлены!");

      handlePhotos();
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

  const { mutate: mutateAddPhotos, isPending: isPendingAddPhoto } = useMutation(
    {
      mutationFn: addRecipePhotos,
      onError: (error: AxiosError<any>) => {
        showErrorNotification(error);
      },
    }
  );

  useEffect(() => {
    const getSteps = async () => {
      const steps: RecipeStepObj[] = [];

      for (const step of data?.data ?? []) {
        const img = await urlToObject(
          getRecipeStepImg(step),
          `${step.step}.webp`
        );

        steps.push({ description: step.description, img });
      }

      return steps;
    };

    getSteps().then((steps) => {
      form.setValues({ formSteps: steps });
    });
  }, [data?.data]);

  const handleSubmit = form.onSubmit((values) => {
    mutateAdd({
      uid,
      steps: values.formSteps.map((s, ind) => ({
        recipeUid: uid,
        description: s.description,
        step: ind + 1,
      })),
    });
  });

  const handleAddStep = () => {
    form.insertListItem("formSteps", {
      description: "",
      img: null,
    });
  };

  const handlePhotos = async () => {
    const formData = new FormData();
    formData.set("uid", uid);

    const { formSteps } = form.getValues();

    for (let ind = 0; ind < formSteps.length; ind++) {
      const step = formSteps[ind];

      const processed = await processImgFile(step.img!);

      formData.set("file", renameFile(processed, `${ind + 1}.webp`));
    }

    mutateAddPhotos(formData);
  };

  const handleRemoveItem = (ind: number) => {
    form.removeListItem("formSteps", ind);

    mutateDeletePhoto({ uid, photos: [`${ind}.webp`] });
  };

  // const handlePhotos = () => {

  // };

  const isPending =
    isPendingDeletePhoto || isPendingAdding || isPendingAddPhoto;

  return (
    <form onSubmit={handleSubmit}>
      <Stack gap={16}>
        <Stack gap={12}>
          {form.getValues().formSteps.map((step, ind) => (
            <Paper radius={12} bg="bg.0" key={ind}>
              <Group align="flex-start" p={12} wrap="nowrap" gap={10}>
                <Stack w="100%" gap={12}>
                  <Textarea
                    size="md"
                    label="Описание"
                    withAsterisk
                    required
                    autosize
                    disabled={isPending}
                    maxRows={5}
                    placeholder="Введите описание"
                    key={form.key(`formSteps.${ind}.description`)}
                    {...form.getInputProps(`formSteps.${ind}.description`)}
                  />

                  <FileInput
                    size="md"
                    label="Изображение"
                    withAsterisk
                    accept="image/webp"
                    required
                    disabled={isPending}
                    placeholder="Выберите изображение"
                    clearable
                    key={form.key(`formSteps.${ind}.img`)}
                    {...form.getInputProps(`formSteps.${ind}.img`)}
                    onChange={(payload) =>
                      form.setFieldValue(`formSteps.${ind}.img`, payload)
                    }
                  />

                  {step.img && (
                    <Image
                      style={{ width: 100, height: 100 }}
                      key={step.img.name}
                      src={URL.createObjectURL(step.img)}
                    />
                  )}
                </Stack>

                <ActionIcon
                  color="accent.0"
                  onClick={() => handleRemoveItem(ind)}
                >
                  <IconTrash size={16} />
                </ActionIcon>
              </Group>
            </Paper>
          ))}

          <Button w="100%" onClick={handleAddStep} variant="accent-reverse">
            Добавить шаг
          </Button>
        </Stack>

        <Group wrap="nowrap" mt={4} justify="space-between">
          <Button disabled={isPending} w="100%" type="submit" variant="accent">
            Сохранить
          </Button>

          <Button w="100%" onClick={onClose} variant="accent-reverse">
            Закрыть
          </Button>
        </Group>
      </Stack>
    </form>
  );
};
