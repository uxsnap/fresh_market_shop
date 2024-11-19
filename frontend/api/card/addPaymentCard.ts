import client from "../client";

type Body = {
  number: string;
  expired: string;
  cvv: string;
};
export const addPaymentCard = (body: Body) => {
  return client.post("/payments/cards", body);
};

addPaymentCard.queryKey = "addPaymentCard";
