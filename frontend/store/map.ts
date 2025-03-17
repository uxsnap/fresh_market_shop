import { MapForm, MapFormInnerFields } from "@/components/Map/context";
import { DEFAULT_MAP_ZOOM } from "@/constants";
import { DeliveryAddress, MapAddress } from "@/types";
import { create, StoreApi, UseBoundStore } from "zustand";
import { immer } from "zustand/middleware/immer";

type MapState = {
  city: string;
  setCity: (v: string) => void;

  mapInstance: ymaps.Map | null;
  setMapInstance: (m: ymaps.Map) => void;

  isMapOpen: boolean;
  setIsMapOpen: (v: boolean) => void;

  searchValue: string;
  setSearchValue: (v: string) => void;

  isFieldsModalOpen: boolean;
  setIsFieldsModalOpen: (v: boolean) => void;

  mapAddress?: MapAddress;
  setMapAddress: (v?: MapAddress) => void;

  deliveryAddress?: DeliveryAddress;
  setDeliveryAddress: (v?: DeliveryAddress) => void;

  handleCenterMove: (coords: number[]) => void;

  fields: MapFormInnerFields;
  setFields: (v: MapFormInnerFields) => void;
  getFields: () => MapFormInnerFields;
};

export const useMapStore: UseBoundStore<StoreApi<MapState>> = create(
  immer((set, get) => ({
    mapInstance: null,
    setMapInstance: (m: ymaps.Map) => set({ mapInstance: m }),

    isFieldsModalOpen: false,
    setIsFieldsModalOpen: (v: boolean) => set({ isFieldsModalOpen: v }),

    searchValue: "",
    setSearchValue: (v: string) => set({ searchValue: v }),

    setMapAddress: (v?: MapAddress) => set({ mapAddress: v }),
    setDeliveryAddress: (v?: DeliveryAddress) => set({ deliveryAddress: v }),

    isMapOpen: false,
    setIsMapOpen: (v: boolean) => set({ isMapOpen: v }),

    handleCenterMove: (coords: number[]) => {
      get().mapInstance?.setCenter(coords, DEFAULT_MAP_ZOOM, {
        duration: 150,
      });
    },

    city: "_",
    setCity: (v: string) => set({ city: v }),

    fields: {},
    setFields: (v: MapFormInnerFields) => set({ fields: v }),

    getFields: () => get().fields,
  }))
);
