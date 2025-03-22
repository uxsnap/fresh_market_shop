import client from "../client";

type Body = {
  // userUid: string;
  orderUid: string;
  cardUid: string;
  sum: number;
  currency: string;
};

export const makePayment = (body: Body) => {
  return client.post("/payments", body);
};

makePayment.queryKey = "makePayment";
