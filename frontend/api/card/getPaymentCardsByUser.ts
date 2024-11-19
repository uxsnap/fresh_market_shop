import { CreditCard } from "@/types";
import client from "../client";

export const getPaymentCardsByUser = () => {
  return client.get<CreditCard[]>("/payments/cards/by_user");
};

getPaymentCardsByUser.queryKey = "getPaymentCardsByUser";
