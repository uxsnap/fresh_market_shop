import { COOKIE_AUTH_TOKENS_NAME } from "@/constants";
import { NextServerResult } from "@/types";
import { cookies } from "next/headers";

export const getAppCookie = async () => {
  const cookieStore = await cookies();

  return cookieStore.get(COOKIE_AUTH_TOKENS_NAME) ?? "";
};

export const getAuthCookieTokensFromServer =
  async (): Promise<NextServerResult> => {
    const result: NextServerResult = { success: false };

    try {
      const appCookie = await getAppCookie();

      result.tokens = appCookie ? JSON.parse(appCookie.value) : undefined;
      result.success = !!result.tokens;
    } catch (e: any) {
      result.success = false;
      result.error = e;
    }

    return result;
  };
