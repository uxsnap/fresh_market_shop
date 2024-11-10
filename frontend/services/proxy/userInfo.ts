import axios from "axios";
import { publicApiErrorResponse } from "@/utils";
import { NextRequest } from "next/server";
import {
  deleteAuthCookies,
  getAuthCookieTokensFromServer,
  parseJwt,
} from "./cookies";

export const proxyUserInfo = async () => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens?.access_jwt) {
      return deleteAuthCookies();
    }

    const parsed = parseJwt(tokens.access_jwt);

    return Response.json(parsed, { status: 200 });
  } catch (error) {
    return publicApiErrorResponse(error);
  }
};
