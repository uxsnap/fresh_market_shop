import { CreditCard } from "@/types";
import { create } from "zustand";

type OrderState = {
  creditCard?: CreditCard;
  setCreditCard: (c: CreditCard) => void;
};

export const useOrderStore = create<OrderState>((set) => ({
  setCreditCard: (c: CreditCard) => set({ creditCard: c }),
}));
