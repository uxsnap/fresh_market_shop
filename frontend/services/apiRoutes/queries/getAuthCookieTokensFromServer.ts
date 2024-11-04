import { COOKIE_AUTH_TOKENS_NAME } from "@/constants";
import { NextServerResult } from "@/types";
import { cookies } from "next/headers";

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
