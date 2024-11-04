import { AuthType } from "@/types";
import { create } from "zustand";

type AuthState = {
  logged?: boolean;
  modalOpen: AuthType | "";
  setLogged: (val?: boolean) => void;
  setModalOpen: (val: AuthType | "") => void;
};

export const useAuthStore = create<AuthState>((set) => ({
  logged: undefined,
  modalOpen: "",
  setModalOpen: (v: AuthType | "") => set({ modalOpen: v }),
  setLogged: (v?: boolean) => set({ logged: v }),
}));
