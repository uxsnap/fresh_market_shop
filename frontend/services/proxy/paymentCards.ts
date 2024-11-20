import axios from "axios";
import {
  deleteAuthCookies,
  getAuthCookieTokensFromServer,
  parseJwt,
} from "./cookies";
import { publicApiErrorResponse } from "@/utils";
import { NextRequest } from "next/server";

export const proxyGetPaymentCardsByUser = async () => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens?.access_jwt) {
      return deleteAuthCookies();
    }

    const parsed = parseJwt(tokens.access_jwt);

    const response = await axios.get(
      `${process.env.NEXT_PUBLIC_API}/payments/cards/by_user/${parsed?.user_uid}`,
      {
        headers: { Authorization: `Bearer ${tokens.access_jwt}` },
      }
    );

    return Response.json(response.data, { status: 200 });
  } catch (error) {
    return publicApiErrorResponse(error);
  }
};

export const proxyAddPaymentCard = async (req: NextRequest) => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens?.access_jwt) {
      return deleteAuthCookies();
    }

    const parsed = parseJwt(tokens.access_jwt);

    const body = await req.json();

    const response = await axios.post(
      `${process.env.NEXT_PUBLIC_API}/payments/cards`,
      { ...body, userUid: parsed?.user_uid },
      {
        headers: { Authorization: `Bearer ${tokens.access_jwt}` },
      }
    );

    return Response.json(response.data, { status: 200 });
  } catch (error) {
    return publicApiErrorResponse(error);
  }
};
