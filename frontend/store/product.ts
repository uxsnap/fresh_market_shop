import { ProductItem } from "@/types";
import { create } from "zustand";

type ProductState = {
  curItem?: ProductItem;
  setCurItem: (c?: ProductItem) => void;
};

export const useProductStore = create<ProductState>((set) => ({
  setCurItem: (c?: ProductItem) => set({ curItem: c }),
}));
