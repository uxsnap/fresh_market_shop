import axios from "axios";
import { getBase64Img, publicApiErrorResponse } from "@/utils";
import {
  deleteAuthCookies,
  getAuthCookieTokensFromServer,
  parseJwt,
} from "./cookies";
import { NextRequest } from "next/server";

export const proxyUpdatePhoto = async (req: NextRequest) => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens?.access_jwt) {
      return deleteAuthCookies();
    }

    const parsed = parseJwt(tokens.access_jwt);

    const body = await req.formData();

    await axios.post(
      `${process.env.NEXT_PUBLIC_API}/users/${parsed?.user_uid}/photo`,
      body,
      {
        headers: { Authorization: `Bearer ${tokens.access_jwt}` },
      }
    );

    return Response.json("", { status: 200 });
  } catch (error) {
    return publicApiErrorResponse(error);
  }
};

export const proxyGetPhoto = async () => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens?.access_jwt) {
      return deleteAuthCookies();
    }

    const parsed = parseJwt(tokens.access_jwt);

    const response = await axios.get(
      `${process.env.NEXT_PUBLIC_API}/assets/photo/${parsed?.user_uid}.webp`,
      {
        headers: { Authorization: `Bearer ${tokens.access_jwt}` },
        responseType: "arraybuffer",
      }
    );

    return Response.json(
      {
        src: getBase64Img(response),
      },
      { status: 200 }
    );
  } catch (error) {
    return publicApiErrorResponse(error);
  }
};
