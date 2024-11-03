import { setAuthCookieTokensFromServer } from "@/services/apiRoutes/authCookie";
import { deleteAuthCookieTokensFromServer } from "@/services/apiRoutes/queries/deleteAuthCookieTokensFromServer";
import { getAuthCookieTokensFromServer } from "@/services/apiRoutes/queries/getAuthCookieTokensFromServer";
import axios from "axios";
import { NextRequest, NextResponse } from "next/server";

export async function GET(req: NextRequest, res: NextResponse) {
  const url = req.url.replace(process.env.NEXT_PUBLIC_API + "", "");

  switch (url) {
    case "/set_auth_tokens":
      const resultSetCookie = setAuthCookieTokensFromServer(req.body, res);

      return NextResponse.json(resultSetCookie, {
        status: resultSetCookie.success ? 200 : 500,
      });

    // return res
    //   .status(resultSetCookie.success ? 200 : 500)
    //   .json(resultSetCookie);
    // case "/delete_auth_tokens":
    //   const resultDeleteTokens = deleteAuthCookieTokensFromServer(res);

    //   return res
    //     .status(resultDeleteTokens.success ? 200 : 500)
    //     .json(resultDeleteTokens);
    default:
      try {
        const headers: Record<string, string>[] = [];

        const { tokens } = await getAuthCookieTokensFromServer();

        if (!(tokens && req.url === "/login")) {
          const response = await axios.request({
            ...req,
            url,
            baseURL: process.env.NEXT_PUBLIC_API,
            data: req.body,
            headers: {
              ...Object.assign({}, ...headers),
            },
          });

          return NextResponse.json(response.data, { status: response.status });
        }

        const { refresh_token, access_token, exp } = tokens;

        if (!isAccessTokenAlmostExpired(exp)) {
          headers.push({
            Authorization: `Bearer ${access_token}`,
          });
        } else {
          try {
            const refreshTokenResponse = await axios.post(
              `${process.env.NEXT_PUBLIC_API}/refresh`,
              {
                refresh_token,
              },
              {
                headers: {
                  Authorization: `Bearer ${access_token}`,
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
      } catch (error) {
        return NextResponse.json(
          axios.isAxiosError(error) ? error?.response?.data : error,
          {
            status: axios.isAxiosError(error)
              ? error?.response?.status || 500
              : 500,
          }
        );
      }
  }
}
