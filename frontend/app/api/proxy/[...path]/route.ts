import { setAuthCookieTokensFromServer } from "@/services/apiRoutes/authCookie";
import { deleteAuthCookieTokensFromServer } from "@/services/apiRoutes/queries/deleteAuthCookieTokensFromServer";
import {
  getAppCookie,
  getAuthCookieTokensFromServer,
} from "@/services/apiRoutes/queries/getAuthCookieTokensFromServer";
import { isAccessTokenAlmostExpired, publicApiErrorResponse } from "@/utils";
import axios from "axios";
import { NextRequest, NextResponse } from "next/server";

export async function POST(req: NextRequest, res: NextResponse) {
  const remoteServer = process.env.NEXT_PUBLIC_API;

  const url = req.url?.replace(
    process.env.NEXT_PUBLIC_API_PROXY_BASE_URL + "",
    ""
  );

  const { tokens } = await getAuthCookieTokensFromServer();

  switch (url) {
    case "/verify":
      if (!tokens) {
        return Response.json({ isValid: false });
      }

      try {
        const response = await axios.request({
          ...req,
          method: "POST",
          url: "/auth/verify/jwt",
          baseURL: process.env.NEXT_PUBLIC_API,
          data: {
            jwt: tokens.refresh_token,
          },
          headers: {},
        });

        return Response.json({ isValid: response.data.isValid });
      } catch (err) {
        console.log(err);

        return Response.json({ isValid: false });
      }
    default:
      try {
        const headers: Record<string, string>[] = [];

        if (tokens) {
          const { refresh_token, access_token, exp } = tokens;

          if (!isAccessTokenAlmostExpired(exp)) {
            headers.push({
              Authorization: `Bearer ${access_token}`,
            });
          } else {
            try {
              const refreshTokenResponse = await axios.post(
                `${remoteServer}/refresh`,
                undefined,
                {
                  headers: {
                    Cookie: `refresh_jwt=${refresh_token}`,
                  },
                }
              );

              headers.push({
                Authorization: `Bearer ${refreshTokenResponse.data.access_token}`,
              });

              setAuthCookieTokensFromServer(refreshTokenResponse.data, res);
            } catch (e) {
              deleteAuthCookieTokensFromServer(res);
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

        return Response.json(response.data, { status: response.status });
      } catch (error) {
        return publicApiErrorResponse(error);
      }
  }
}
