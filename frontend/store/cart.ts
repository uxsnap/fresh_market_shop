import { CartItem, DeliveryData, ProductItem } from "@/types";
import { immer } from "zustand/middleware/immer";
import { create, StoreApi, UseBoundStore } from "zustand";

type CartState = {
  items: Record<string, CartItem>;
  incCartItem: (itemId: string) => void;
  decCartItem: (itemId: string) => void;
  addCartItem: (product: ProductItem) => void;
  removeCartItem: (itemId: string) => void;
  removeAllItems: () => void;
  getFullPrice: () => number;
  getItemsPrice: () => number;
  getCount: (itemId: string) => number;
  delivery?: DeliveryData;
  setDelivery: (delivery: DeliveryData) => void;
};

export const useCartStore: UseBoundStore<StoreApi<CartState>> = create(
  immer((set, get) => ({
    items: {},
    incCartItem: (itemId: string) => {
      return set((state) => {
        const item = state.items[itemId];

        if (!item) {
          return state;
        }

        item!.count += 1;

        return state;
      });
    },
    decCartItem: (itemId: string) => {
      return set((state) => {
        const item = state.items[itemId];

        if (!item) {
          return state;
        }

        if (item.count === 1) {
          delete state.items[itemId];

          return state;
        }

        item!.count -= 1;

        return state;
      });
    },
    addCartItem: (product: ProductItem) => {
      return set((state) => {
        state.items[product.id] = { product, count: 1 };

        return state;
      });
    },
    removeCartItem: (itemId: string) => {
      return set((state) => {
        delete state.items[itemId];

        return state;
      });
    },
    getItemsPrice() {
      const arr = Object.values(get().items);

      if (!arr.length) {
        return 0;
      }

      return arr.reduce(
        (acc, item) => acc + item.product.price * item.count,
        0
      );
    },
    getFullPrice() {
      return get().getItemsPrice() + (get().delivery?.price ?? 0);
    },
    getCount(itemId: string) {
      return get().items[itemId]?.count ?? 0;
    },

    setDelivery: (delivery: DeliveryData) => {
      return set((state) => {
        state.delivery = delivery;

        return state;
      });
    },
    removeAllItems() {
      return set((state) => {
        state.items = {};

        return state;
      });
    },
  }))
);
