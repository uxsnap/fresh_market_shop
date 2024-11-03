import { COOKIE_AUTH_TOKENS_NAME } from "@/constants";
import { NextServerResult } from "@/types";
import cookie from "cookie";
import { NextResponse } from "next/server";

export const deleteAuthCookieTokensFromServer = (res: NextResponse) => {
  const result: NextServerResult = { success: true };
  try {
    res.headers.set(
      "Set-Cookie",
      cookie.serialize(COOKIE_AUTH_TOKENS_NAME, "", {
        httpOnly: true,
        secure: process.env.NODE_ENV !== "development",
        maxAge: -1,
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
