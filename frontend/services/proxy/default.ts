import axios from "axios";
import {
  deleteAuthCookies,
  getAuthCookieTokensFromServer,
  isAccessTokenAlmostExpired,
  parseResponseCookies,
} from "./cookies";
import { publicApiErrorResponse } from "@/utils";
import { NextRequest } from "next/server";
import { cookies } from "next/headers";

export const proxyDefault = async (req: NextRequest) => {
  const url = req.url?.replace(
    process.env.NEXT_PUBLIC_API_PROXY_BASE_URL + "",
    ""
  );

  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    const headers: Record<string, string>[] = [];

    if (tokens?.refresh_jwt && tokens.access_jwt) {
      const { refresh_jwt, access_jwt } = tokens;

      if (!isAccessTokenAlmostExpired(access_jwt)) {
        headers.push({
          Authorization: `Bearer ${access_jwt}`,
        });
      } else {
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

          const parsed =
            parseResponseCookies(refreshTokenResponse)["access_jwt"];

          const cookieStore = await cookies();

          cookieStore.set("access_jwt", parsed);
        } catch (e) {
          return deleteAuthCookies();
        }
      }
    }

    const response = await axios.request({
      ...req,
      url,
      baseURL: process.env.NEXT_PUBLIC_API,
      data: req.body,
      headers: {
        ...Object.assign({}, ...headers),
      },
    });

    return Response.json(response.data, {
      status: response.status,
    });
  } catch (error) {
    return publicApiErrorResponse(error);
  }
};
