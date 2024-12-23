import { AdminTab } from "@/types";
import { immer } from "zustand/middleware/immer";
import { create, StoreApi, UseBoundStore } from "zustand";

type AdminState = {
  tab: AdminTab;
  setTab: (newTab: AdminTab) => void;
};

export const useAdminStore: UseBoundStore<StoreApi<AdminState>> = create(
  immer((set) => ({
    tab: AdminTab.admins,
    setTab: (newTab) => {
      return set((state) => {
        state.tab = newTab;

        return state;
      });
    },
  }))
);
