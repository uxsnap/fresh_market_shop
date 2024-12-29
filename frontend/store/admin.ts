import { AdminTab, ProductItem } from "@/types";
import { create } from "zustand";
import { immer } from "zustand/middleware/immer";

type AdminState = {
  modalOpen?: boolean;
  productItem?: ProductItem;
  setModalOpen: (val: boolean) => void;
  setProductItem: (val?: ProductItem) => void;
  tab: AdminTab;
  setTab: (newTab: AdminTab) => void;
};

export const useAdminStore = create<AdminState>()(
  immer((set) => ({
    modalOpen: false,
    tab: AdminTab.admins,
    setTab: (newTab: AdminTab) => {
      return set((state) => {
        state.tab = newTab;

        return state;
      });
    },
    setModalOpen: (val: boolean) => {
      return set((state) => {
        state.modalOpen = val;
        return state;
      });
    },
    setProductItem: (val) => {
      return set((state) => {
        state.productItem = val;
        return state;
      });
    },
  }))
);
