import {
  ErrorWrapper,
  ProductItem,
  ProductWithPhotos,
  RecipeStep,
} from "@/types";
import { notifications } from "@mantine/notifications";
import { AxiosError } from "axios";

import dayjs from "dayjs";
import duration from "dayjs/plugin/duration";

export const dayJs = (() => {
  dayjs.extend(duration);

  return dayjs;
})();

export const getFallbackImg = (name: string) => {
  return "https://placehold.co/800x600/FFF/4F463D?font=roboto&text=" + name;
};

export const convertProductToProductItem = (
  item: ProductWithPhotos
): ProductItem => ({
  id: item.product.uid,
  imgs: (item.photos ?? []).map(
    (p) => `${process.env.NEXT_PUBLIC_API}/${p.path}`
  ),
  price: item.product.price,
  name: item.product.name,
  weight: item.product.weight,
  ccal: item.product.ccal,
  description: item.product.description,
});

export const formatDuration = (duration: number) => {
  const hours = dayJs.duration(duration / 1000).hours();
  const minutes = dayJs.duration(duration / 1000).minutes();

  return `${hours ? hours + "ч " : ""}${minutes ? minutes + "м" : ""}`;
};

export const getRecipeBg = (uid: string) => {
  return `${process.env.NEXT_PUBLIC_API}/assets/recipes/${uid}/main.jpg`;
};

export const getRecipeStepImg = (step: RecipeStep) => {
  return `${process.env.NEXT_PUBLIC_API}/assets/recipes/${step.recipeUid}/${step.step}.jpg`;
};

export const publicApiErrorResponse = (error: unknown) => {
  return Response.json((error as any)?.response.data, {
    status: 500,
  });
};

export const getErrorBody = (error: AxiosError<{ error: ErrorWrapper }>) => {
  return error.response?.data?.error;
};

export const showErrorNotification = (
  error: AxiosError<{ error: ErrorWrapper }>
) => {
  notifications.show({
    title: "Ошибка!",
    message: getErrorBody(error)?.message,
    color: "red",
  });
};

export const showSuccessNotification = (message: string) => {
  notifications.show({
    title: "Успешно!",
    message,
    color: "secondary.0",
  });
};

export const isDateNull = (date?: string) => {
  const d = dayJs(date);

  return !date || !d.isValid() || d.year() <= 1;
};

export * from "./img";
