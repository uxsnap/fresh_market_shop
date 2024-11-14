import { AuthType, JwtData } from "@/types";
import { create } from "zustand";

type AuthState = {
  logged?: boolean;
  modalOpen: AuthType | "";
  setLogged: (val?: boolean) => void;
  setModalOpen: (val: AuthType | "") => void;
  userInfo?: JwtData;
  setUserInfo: (val?: JwtData) => void;
};

export const useAuthStore = create<AuthState>((set) => ({
  userInfo: undefined,
  logged: undefined,
  modalOpen: "",
  setModalOpen: (v: AuthType | "") => set({ modalOpen: v }),
  setLogged: (v?: boolean) => set({ logged: v }),
  setUserInfo: (v?: JwtData) => set({ userInfo: v }),
}));
