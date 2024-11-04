import { publicApiErrorResponse } from "@/utils";
import axios from "axios";
import { NextRequest } from "next/server";
import { deleteAuthCookies, getAuthCookieTokensFromServer } from "./cookies";

export const proxyLogout = async (req: NextRequest) => {
  try {
    const body = await req.json();

    const { tokens } = await getAuthCookieTokensFromServer();

    await axios.post(`${process.env.NEXT_PUBLIC_API}/auth/logout`, body, {
      headers: { Bearer: `${tokens?.access_jwt}` },
    });

    return deleteAuthCookies();
  } catch (error) {
    return publicApiErrorResponse(error);
  }
};
