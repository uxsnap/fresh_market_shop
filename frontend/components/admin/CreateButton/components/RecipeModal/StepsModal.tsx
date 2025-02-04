import { ImgsUpload } from "@/components/admin/ImgsUpload";
import { RecipeStepObj } from "@/types";
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
import { useEffect } from "react";

export type Props = {
  steps: RecipeStepObj[];
  onChange: (s: RecipeStepObj[]) => void;
  onClose: () => void;
};

export const StepsModal = ({ steps, onChange, onClose }: Props) => {
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

  useEffect(() => {
    form.setValues({ formSteps: steps });
  }, [steps]);

  const handleSubmit = form.onSubmit((values) => {
    onChange(values.formSteps);
  });

  const handleAddStep = () => {
    form.insertListItem("formSteps", {
      description: "",
      img: null,
    });
  };

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
                      key={step.img.name}
                      src={URL.createObjectURL(step.img)}
                    />
                  )}
                </Stack>

                <ActionIcon
                  color="accent.0"
                  onClick={() => form.removeListItem("formSteps", ind)}
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
          <Button w="100%" type="submit" variant="accent">
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
