import { UserAddress } from "@/types";
import { YMapsApi } from "@pbe/react-yandex-maps/typings/util/typing";
import { create } from "zustand";

type MapState = {
  isMapOpen: boolean;
  setIsMapOpen: (v: boolean) => void;
  map: YMapsApi | null;
  setMap: (m: YMapsApi) => void;
  searchValue: string;
  setSearchValue: (v: string) => void;
  isFieldsModalOpen: boolean;
  setIsFieldsModalOpen: (v: boolean) => void;
  activeAddress?: UserAddress;
  setActiveAddress: (v: UserAddress) => void;
};

export const useMapStore = create<MapState>((set) => ({
  map: null,
  setMap: (m: YMapsApi) => set({ map: m }),
  isFieldsModalOpen: false,
  setIsFieldsModalOpen: (v: boolean) => set({ isFieldsModalOpen: v }),
  searchValue: "",
  setSearchValue: (v: string) => set({ searchValue: v }),
  setActiveAddress: (v: UserAddress) => set({ activeAddress: v }),
  isMapOpen: false,
  setIsMapOpen: (v: boolean) => set({ isMapOpen: v }),
}));
