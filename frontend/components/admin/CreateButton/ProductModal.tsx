import { getCategories } from "@/api/categories/getCategories";
import { createProduct } from "@/api/products/createProduct";
import { getProducts } from "@/api/products/getProducts";
import { useAdminStore } from "@/store/admin";
import { showErrorNotification, showSuccessNotification } from "@/utils";
import {
  Button,
  Group,
  NumberInput,
  Select,
  Stack,
  Textarea,
  TextInput,
} from "@mantine/core";
import { hasLength, isNotEmpty, useForm } from "@mantine/form";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { useEffect } from "react";

type Props = {
  onClose: () => void;
};

const weightData = [
  { label: "20 грамм", value: "20" },
  { label: "100 грамм", value: "100" },
];

export const ProductModal = ({ onClose }: Props) => {
  const productItem = useAdminStore((s) => s.productItem);
  const setProductItem = useAdminStore((s) => s.setProductItem);

  const queryClient = useQueryClient();

  const form = useForm({
    mode: "uncontrolled",
    initialValues: {
      name: "",
      description: "",
      categoryUid: "",
      ccal: 0,
      price: 0,
      weight: "",
    },
    validate: {
      name: hasLength({ min: 1 }, "Длина названия должна быть больше 1"),
      description: hasLength({ min: 1 }, "Длина описания должна быть больше 1"),
      categoryUid: isNotEmpty("Категория не должна быть пустой"),
      ccal: isNotEmpty("Калорийность не должна быть пустой"),
      price: isNotEmpty("Цена не должна быть пустой"),
      weight: isNotEmpty("Вес не должна быть пустой"),
    },
  });

  useEffect(() => {
    if (!productItem) {
      return;
    }

    form.setValues({
      name: productItem.name,
      description: productItem.description,
      ccal: productItem.ccal,
      price: productItem.price,
      weight: productItem.weight + "",
      categoryUid: productItem.categoryUid,
    });
  }, [productItem]);

  const handleClose = () => {
    setProductItem(undefined);
    onClose();
  };

  const { mutate, isPending } = useMutation({
    mutationFn: createProduct,
    mutationKey: [createProduct.queryKey],
    onSuccess: () => {
      onClose();

      queryClient.invalidateQueries({
        queryKey: [getProducts.queryKey],
      });

      showSuccessNotification("Продукт успешно добавлен!");
    },
    onError: (error: AxiosError<any>) => {
      showErrorNotification(error);
    },
  });

  const handleSubmit = form.onSubmit((values) => {
    mutate({
      ...values,
      weight: parseInt(values.weight, 10),
    });
  });

  const { data: categories } = useQuery({
    queryKey: [getCategories.queryKey],
    queryFn: getCategories,
    refetchOnWindowFocus: false,
    staleTime: Infinity,
    select(data) {
      return data.data.map((c) => ({
        label: c.name,
        value: c.uid,
      }));
    },
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

        <Select
          w="100%"
          size="md"
          label="Категория"
          placeholder="Выберите категорию"
          data={categories ?? []}
          withAsterisk
          required
          withScrollArea={false}
          styles={{ dropdown: { maxHeight: 130, overflowY: "auto" } }}
          key={form.key("categoryUid")}
          {...form.getInputProps("categoryUid")}
          comboboxProps={{ withinPortal: false }}
        />

        <Textarea
          radius="md"
          label="Описание"
          placeholder="Введите описание"
          withAsterisk
          required
          {...form.getInputProps("description")}
          resize="vertical"
          minRows={10}
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

        <NumberInput
          w="100%"
          min={1}
          hideControls
          allowLeadingZeros={false}
          allowNegative={false}
          allowDecimal={false}
          withAsterisk
          lh={1}
          required
          size="md"
          label="Цена"
          placeholder="Введите цену"
          key={form.key("price")}
          {...form.getInputProps("price")}
        />

        <Select
          w="100%"
          size="md"
          label="Вес"
          placeholder="Выберите вес"
          data={weightData}
          withAsterisk
          required
          withScrollArea={false}
          styles={{ dropdown: { maxHeight: 130, overflowY: "auto" } }}
          key={form.key("weight")}
          {...form.getInputProps("weight")}
          comboboxProps={{ withinPortal: false }}
        />

        <Group wrap="nowrap" mt={4} justify="space-between">
          <Button disabled={isPending} w="100%" type="submit" variant="accent">
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
