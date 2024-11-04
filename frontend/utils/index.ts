import { JwtData, ProductItem, ProductWithPhotos, RecipeStep } from "@/types";
import axios from "axios";

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
  item: ProductWithPhotos,
  _: any
): ProductItem => ({
  id: item.product.uid,
  imgs: (item.photos ?? []).map(
    (p) => `${process.env.NEXT_PUBLIC_API}/${p.path}`
  ),
  price: item.product.price,
  name: item.product.name,
  weight: item.product.weight,
  info: `${item.product.weight}грамм/${item.product.ccal}ккал`,
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

export const parseJwt = (token?: string): JwtData | undefined => {
  if (!token) {
    return;
  }

  const base64Url = token.split(".")[1];
  const base64 = base64Url.replace("-", "+").replace("_", "/");
  return JSON.parse(Buffer.from(base64, "base64").toString());
};

export const isAccessTokenAlmostExpired = (token: string) => {
  const parsed = parseJwt(token);

  if (!parsed) {
    return true;
  }

  const expiredTime = parseInt(parsed.expired_at, 10);

  const SECONDS_TO_EXPIRE = 15000;

  return Date.now() + SECONDS_TO_EXPIRE >= expiredTime * 1000;
};

export const publicApiErrorResponse = (error: unknown) => {
  return Response.json(axios.isAxiosError(error) ? error?.response : error, {
    status: axios.isAxiosError(error) ? error?.response?.status || 500 : 500,
  });
};
