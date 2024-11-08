import {
  DAY_IN_SECONDS,
  ONE_HOUR_IN_SECONDS,
  SECONDS_TO_EXPIRE,
} from "@/constants";
import { JwtData, NextServerResult } from "@/types";
import { AxiosResponse } from "axios";
import { cookies } from "next/headers";

const cookie = require("cookie");

export const getAppCookie = async () => {
  const cookieStore = await cookies();

  return {
    access_jwt: cookieStore.get("access_jwt")?.value,
    refresh_jwt: cookieStore.get("refresh_jwt")?.value,
  };
};

export const getAuthCookieTokensFromServer =
  async (): Promise<NextServerResult> => {
    const result: NextServerResult = { success: false };

    try {
      const { access_jwt, refresh_jwt } = await getAppCookie();

      if (access_jwt && refresh_jwt) {
        result.tokens = { access_jwt, refresh_jwt };
      }

      result.success = !!result.tokens;
    } catch (e: any) {
      result.success = false;
      result.error = e;
    }

    return result;
  };

export const parseResponseCookies = (res: AxiosResponse<any, any>) => {
  return res.headers["set-cookie"]!.reduce(
    (acc, c) => {
      const p = cookie.parse(c);

      return { ...acc, ...p };
    },
    {} as Record<string, string>
  );
};

export const serializeCookie = (name: string, val: string, age: number) => {
  return cookie.serialize(name, val, {
    httpOnly: true,
    secure: process.env.NODE_ENV !== "development",
    maxAge: age,
    sameSite: "strict",
    path: "/",
  });
};

export const deleteAuthCookies = (resp = {}) => {
  return Response.json(resp, {
    status: 200,
    headers: [
      ["Set-Cookie", serializeCookie("access_jwt", "", -1)],
      ["Set-Cookie", serializeCookie("refresh_jwt", "", -1)],
    ],
  });
};

export const prepareAuthCookies = (
  parsed: Record<string, string>
): HeadersInit => {
  return [
    [
      "Set-Cookie",
      serializeCookie("access_jwt", parsed["access_jwt"], ONE_HOUR_IN_SECONDS),
    ],
    [
      "Set-Cookie",
      serializeCookie("refresh_jwt", parsed["refresh_jwt"], DAY_IN_SECONDS),
    ],
  ];
};

export const setAuthCookiesFromResponse = (
  data: any,
  parsed: Record<string, string>
) =>
  Response.json(data, {
    status: 200,
    headers: prepareAuthCookies(parsed),
  });

export const parseJwt = (token?: string): JwtData | undefined => {
  if (!token) {
    return;
  }

  const base64Url = token.split(".")[1];
  const base64 = base64Url.replace("-", "+").replace("_", "/");
  return JSON.parse(Buffer.from(base64, "base64").toString());
};

export const isAccessTokenAlmostExpired = (token?: string) => {
  if (!token) {
    return true;
  }

  const parsed = parseJwt(token);

  if (!parsed) {
    return true;
  }

  const expiredTime = parseInt(parsed.expired_at, 10);

  return Date.now() + SECONDS_TO_EXPIRE >= expiredTime * 1000;
};
