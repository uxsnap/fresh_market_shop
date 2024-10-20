import { Product, ProductWithPhotos } from "@/types";

export const getFallbackImg = (name: string) => {
  return "https://placehold.co/800x600/FFF/4F463D?font=roboto&text=" + name;
};

export const convertProductToProductItem = (item: ProductWithPhotos, _: any) => ({
  imgs: (item.photos ?? []).map(
    (p) => `${process.env.NEXT_PUBLIC_API}/${p.path}`
  ),
  price: item.product.price,
  name: item.product.name,
  info: `${item.product.weight}грамм/${item.product.ccal}ккал`,
});
