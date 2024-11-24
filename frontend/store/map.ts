import { Address, DeliveryAddress, MapAddress } from "@/types";
import { YMapsApi } from "@pbe/react-yandex-maps/typings/util/typing";
import { create } from "zustand";

type MapState = {
  map: YMapsApi | null;
  setMap: (m: YMapsApi) => void;

  isMapOpen: boolean;
  setIsMapOpen: (v: boolean) => void;

  searchValue: string;
  setSearchValue: (v: string) => void;

  isFieldsModalOpen: boolean;
  setIsFieldsModalOpen: (v: boolean) => void;

  mapAddress?: MapAddress;
  setMapAddress: (v?: MapAddress) => void;

  deliveryAddress?: DeliveryAddress;
  setDeliveryAddress: (v: DeliveryAddress) => void;
};

export const useMapStore = create<MapState>((set) => ({
  map: null,
  setMap: (m: YMapsApi) => set({ map: m }),

  isFieldsModalOpen: false,
  setIsFieldsModalOpen: (v: boolean) => set({ isFieldsModalOpen: v }),

  searchValue: "",
  setSearchValue: (v: string) => set({ searchValue: v }),

  setMapAddress: (v?: MapAddress) => set({ mapAddress: v }),
  setDeliveryAddress: (v?: DeliveryAddress) => set({ deliveryAddress: v }),

  isMapOpen: false,
  setIsMapOpen: (v: boolean) => set({ isMapOpen: v }),
}));
