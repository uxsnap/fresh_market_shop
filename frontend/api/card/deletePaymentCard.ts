import client from "../client";

export const deletePaymentCard = (cardId: string) => {
  return client.delete(`/payments/cards/${cardId}`);
};

deletePaymentCard.queryKey = "deletePaymentCard";
