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

export const showErrorNotification = (
  error: AxiosError<{ error: ErrorWrapper }>
) => {
  notifications.show({
    title: "Ошибка!",
    message: error.response?.data.error.message,
    color: "red",
  });
};
