import client from "../client";

type Body = {
  orderUid: string;
  cardUid: string;
  sum: number;
  currency: string;
};

export const makePayment = (body: Body) => {
  return client.post("/payments", body);
};

makePayment.queryKey = "makePayment";
