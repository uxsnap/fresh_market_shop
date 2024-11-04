import { setAuthCookieTokensFromServer } from "@/services/apiRoutes/authCookie";
import { deleteAuthCookieTokensFromServer } from "@/services/apiRoutes/queries/deleteAuthCookieTokensFromServer";
import { getAuthCookieTokensFromServer } from "@/services/apiRoutes/queries/getAuthCookieTokensFromServer";
import {
  isAccessTokenAlmostExpired,
  publicApiErrorResponse,
} from "@/utils";
import axios from "axios";
import { NextRequest, NextResponse } from "next/server";
const cookie = require("cookie");

export async function POST(req: NextRequest, res: NextResponse) {
  const remoteServer = process.env.NEXT_PUBLIC_API;

  const url = req.url?.replace(
    process.env.NEXT_PUBLIC_API_PROXY_BASE_URL + "",
    ""
  );
  switch (url) {
    case "/login":
      try {
        const body = await req.json();
        const response = await axios.post(
          `${process.env.NEXT_PUBLIC_API}/auth/login`,
          body
        );

        const parsed = response.headers["set-cookie"]!.reduce(
          (acc, c) => {
            const p = cookie.parse(c);

            return { ...acc, ...p };
          },
          {} as Record<string, string>
        );

        return Response.json(response.data, {
          status: 200,

          headers: [
            ["Set-Cookie", `access_jwt=${parsed["access_jwt"]}`],
            ["Set-Cookie", `refresh_jwt=${parsed["refresh_jwt"]}`],
          ],
        });
      } catch (error) {
        return publicApiErrorResponse(error);
      }

    case "/verify":
      const { tokens } = await getAuthCookieTokensFromServer();

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
            jwt: tokens.access_jwt,
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
        const { tokens } = await getAuthCookieTokensFromServer();

        const headers: Record<string, string>[] = [];

        if (tokens) {
          const { refresh_jwt, access_jwt } = tokens;

          if (!isAccessTokenAlmostExpired("")) {
            headers.push({
              Authorization: `Bearer ${access_jwt}`,
            });
          } else {
            try {
              const refreshTokenResponse = await axios.post(
                `${remoteServer}/refresh`,
                undefined,
                {
                  headers: {
                    Cookie: `refresh_jwt=${refresh_jwt}`,
                  },
                }
              );

              headers.push({
                Authorization: `Bearer ${refreshTokenResponse.data.access_jwt}`,
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
