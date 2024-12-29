import { NextRequest } from "next/server";
import { deleteAuthCookies, getAuthCookieTokensFromServer } from "./cookies";
import axios from "axios";
import { publicApiErrorResponse } from "@/utils";

export const proxyDeleteProduct = async (req: NextRequest) => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens?.access_jwt) {
      return deleteAuthCookies();
    }

    const body = await req.json();

    await axios.delete(`${process.env.NEXT_PUBLIC_API}/products/${body.uid}`, {
      headers: { Authorization: `Bearer ${tokens.access_jwt}` },
    });

    return Response.json(
      {},
      {
        status: 200,
      }
    );
  } catch (error) {
    return publicApiErrorResponse(error);
  }
};

export const proxyReviveProduct = async (req: NextRequest) => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens?.access_jwt) {
      return deleteAuthCookies();
    }

    const body = await req.json();

    await axios.patch(`${process.env.NEXT_PUBLIC_API}/products/${body.uid}`, {
      headers: { Authorization: `Bearer ${tokens.access_jwt}` },
    });

    return Response.json(
      {},
      {
        status: 200,
      }
    );
  } catch (error) {
    return publicApiErrorResponse(error);
  }
};