import { NextServerResult } from "@/types";
import axios from "axios";
import { deleteAuthCookies, parseJwt, parseResponseCookies } from "./cookies";
import { cookies } from "next/headers";

export const refresh = async (tokens: NextServerResult["tokens"]) => {
  const { refresh_jwt, access_jwt } = tokens!;

  try {
    const refreshTokenResponse = await axios.post(
      `${process.env.NEXT_PUBLIC_API}/auth/refresh`,
      undefined,
      {
        headers: {
          Cookie: `refresh_jwt=${refresh_jwt}`,
        },
      }
    );

    const parsed = parseResponseCookies(refreshTokenResponse)["access_jwt"];

    const cookieStore = await cookies();

    cookieStore.set("access_jwt", parsed);

    return true;
  } catch (e) {
    const parsedJwt = parseJwt(access_jwt);

    await axios.post(
      `${process.env.NEXT_PUBLIC_API}/auth/logout`,
      { uid: parsedJwt!.user_uid },
      {
        headers: { Authorization: `Bearer ${access_jwt}` },
      }
    );

    return deleteAuthCookies();
  }
};
