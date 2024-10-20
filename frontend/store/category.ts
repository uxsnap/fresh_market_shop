import { create } from "zustand";

type CategoryState = {
  curCategory: string;
  setCategory: (c: string) => void;
};

export const useCategoryStore = create<CategoryState>((set) => ({
  curCategory: "",
  setCategory: (c: string) => set({ curCategory: c }),
}));
