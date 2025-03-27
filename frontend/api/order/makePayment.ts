import client from "../client";

type Body = {
  orderUid: string;
  cardUid: string;
  deliveryUid: string;
};

export const makePayment = (body: Body) => {
  return client.post("/orders/pay", body);
};

makePayment.queryKey = "makePayment";
