import { create } from "zustand";

type SearchStore = {
  curName: string;
  setCurName: (c: string) => void;
};

export const useSearchStore = create<SearchStore>((set) => ({
  curName: "",
  setCurName: (c: string) => set({ curName: c }),
}));
