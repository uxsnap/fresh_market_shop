import axios from "axios";
import { publicApiErrorResponse } from "@/utils";
import {
  deleteAuthCookies,
  getAuthCookieTokensFromServer,
  parseJwt,
  parseResponseCookies,
  setAuthCookiesFromResponse,
} from "./cookies";
import { NextRequest } from "next/server";

export const proxyUpdateUser = async (req: NextRequest) => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens?.access_jwt) {
      return deleteAuthCookies();
    }

    const parsed = parseJwt(tokens.access_jwt);

    const body = await req.json();

    const response = await axios.put(
      `${process.env.NEXT_PUBLIC_API}/users`,
      { ...body, uid: parsed?.user_uid },
      {
        headers: { Authorization: `Bearer ${tokens.access_jwt}` },
      }
    );

    const parsedResponse = parseResponseCookies(response);

    return setAuthCookiesFromResponse(response.data, parsedResponse);
  } catch (error) {
    return publicApiErrorResponse(error);
  }
};
