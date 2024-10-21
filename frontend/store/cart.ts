import { CartItem, ProductItem } from "@/types";
import { create } from "zustand";

type CartState = {
  items: Map<string, CartItem>;
  incCartItem: (itemId: string) => void;
  decCartItem: (itemId: string) => void;
  addCartItem: (itemId: string, product: ProductItem) => void;
  removeCartItem: (itemId: string) => void;
};

export const useCartStore = create<CartState>((set) => ({
  // @ts-ignore
  items: new Map<string, CartItem>(),
  incCartItem: (itemId: string) => {
    return set((state) => {
      const item = state.items.get(itemId);

      if (!item) {
        return state;
      }

      item!.count += 1;

      return { ...state };
    });
  },
  decCartItem: (itemId: string) => {
    return set((state) => {
      const item = state.items.get(itemId);

      if (!item || item.count === 0) {
        return state;
      }

      item!.count -= 1;

      return { ...state };
    });
  },
  addCartItem: (itemId: string, product: ProductItem) => {
    return set((state) => {
      state.items.set(itemId, { product, count: 0 });

      return { ...state };
    });
  },
  removeCartItem: (itemId: string) => {
    return set((state) => {
      state.items.delete(itemId);

      return { ...state };
    });
  },
}));
