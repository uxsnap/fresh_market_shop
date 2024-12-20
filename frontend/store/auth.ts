import { AuthType, JwtData } from "@/types";
import { create } from "zustand";

type AuthState = {
  logged?: boolean;
  admin?: boolean;
  modalOpen: AuthType | "";
  setLogged: (val?: boolean) => void;
  setAdmin: (val?: boolean) => void;
  setModalOpen: (val: AuthType | "") => void;
  userInfo?: JwtData;
  setUserInfo: (val?: JwtData) => void;
};

export const useAuthStore = create<AuthState>((set) => ({
  userInfo: undefined,
  logged: undefined,
  admin: undefined,
  modalOpen: "",
  setModalOpen: (v: AuthType | "") => set({ modalOpen: v }),
  setLogged: (v?: boolean) => set({ logged: v }),
  setAdmin: (v?: boolean) => set({ admin: v }),
  setUserInfo: (v?: JwtData) => set({ userInfo: v }),
}));
