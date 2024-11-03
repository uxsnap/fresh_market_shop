import { COOKIE_AUTH_TOKENS_NAME } from "@/constants";
import { NextServerResult } from "@/types";
import cookie from "cookie";
import { IncomingMessage } from "http";
import { NextApiResponse } from "next";
import { NextResponse } from "next/server";

export const parseAuthCookieTokens = (cookieStr: string) => {
  const authCookie = cookie.parse(cookieStr)[COOKIE_AUTH_TOKENS_NAME];

  let data = null;

  if (authCookie) {
    try {
      data = JSON.parse(authCookie);
    } catch (e) {}
  }

  return data;
};

export const getAuthCookieTokens = (req: IncomingMessage) => {
  if (!req.headers.cookie) {
    return null;
  }

  return parseAuthCookieTokens(req.headers.cookie);
};

export const setAuthCookieTokensFromServer = (
  tokenResponse: any,
  res: NextResponse
): NextServerResult => {
  const result: NextServerResult = { success: true };

  try {
    const ONE_YEAR_IN_SECONDS = 365 * 24 * 60 * 60;

    res.headers.set(
      "Set-Cookie",
      cookie.serialize(COOKIE_AUTH_TOKENS_NAME, JSON.stringify(tokenResponse), {
        httpOnly: true,
        secure: process.env.NODE_ENV !== "development",
        maxAge: ONE_YEAR_IN_SECONDS,
        sameSite: "strict",
        path: "/",
      })
    );
  } catch (e: any) {
    result.success = false;
    result.error = e;
  }

  return result;
};
