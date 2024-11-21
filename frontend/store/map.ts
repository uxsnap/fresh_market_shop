import { YMapsApi } from "@pbe/react-yandex-maps/typings/util/typing";
import { create } from "zustand";

type MapState = {
  map: YMapsApi | null;
  setMap: (m: YMapsApi) => void;
  searchValue: string;
  setSearchValue: (v: string) => void;
  isFieldsModalOpen: boolean;
  setIsFieldsModalOpen: (v: boolean) => void;
};

export const useMapStore = create<MapState>((set) => ({
  map: null,
  setMap: (m: YMapsApi) => set({ map: m }),
  isFieldsModalOpen: false,
  setIsFieldsModalOpen: (v: boolean) => set({ isFieldsModalOpen: v }),
  searchValue: "",
  setSearchValue: (v: string) => set({ searchValue: v }),
}));
