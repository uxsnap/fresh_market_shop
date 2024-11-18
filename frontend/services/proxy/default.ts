import axios from "axios";
import {
  getAuthCookieTokensFromServer,
  isAccessTokenAlmostExpired,
} from "./cookies";
import { publicApiErrorResponse } from "@/utils";
import { NextRequest } from "next/server";
import { refresh } from "./refresh";

export const proxyDefault = async (req: NextRequest, body?: any) => {
  const url = req.url?.replace(
    process.env.NEXT_PUBLIC_API_PROXY_BASE_URL + "",
    ""
  );

  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    const headers: Record<string, string>[] = [];

    if (tokens?.refresh_jwt && tokens.access_jwt) {
      const { access_jwt } = tokens;

      if (!isAccessTokenAlmostExpired(access_jwt)) {
        headers.push({
          Authorization: `Bearer ${access_jwt}`,
        });
      } else {
        const res = await refresh(tokens);

        if (res !== true) {
          return res;
        }
      }
    }

    const response = await axios.request({
      ...req,
      url,
      baseURL: process.env.NEXT_PUBLIC_API,
      data: body,
      method: req.method,
      headers: {
        ...Object.assign(
          {
            "Content-Type": "application/json",
          },
          ...headers
        ),
      },
    });

    return Response.json(response.data, {
      status: response.status,
    });
  } catch (error) {
    return publicApiErrorResponse(error);
  }
};
