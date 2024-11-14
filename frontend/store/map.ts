import { YMapsApi } from "@pbe/react-yandex-maps/typings/util/typing";
import { create } from "zustand";

type MapState = {
  map: YMapsApi | null;
  setMap: (m: YMapsApi) => void;
};

export const useMapStore = create<MapState>((set) => ({
  map: null,
  setMap: (m: YMapsApi) => set({ map: m }),
}));
