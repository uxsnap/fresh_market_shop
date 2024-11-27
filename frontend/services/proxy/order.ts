import { publicApiErrorResponse } from "@/utils";
import {
  deleteAuthCookies,
  getAuthCookieTokensFromServer,
  parseJwt,
} from "./cookies";
import axios from "axios";

export const proxyOrderHistory = async () => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens?.access_jwt) {
      return deleteAuthCookies([]);
    }

    const parsed = parseJwt(tokens.access_jwt);

    const response = await axios.get(
      `${process.env.NEXT_PUBLIC_API}/orders/${parsed?.user_uid}/history`,
      {
        headers: { Authorization: `Bearer ${tokens.access_jwt}` },
      }
    );

    return Response.json(response.data, { status: 200 });
  } catch (error) {
    return publicApiErrorResponse(error);
  }
};
