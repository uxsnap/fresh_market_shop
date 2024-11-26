import axios from "axios";
import { publicApiErrorResponse } from "@/utils";
import {
  deleteAuthCookies,
  getAuthCookieTokensFromServer,
  parseJwt,
} from "./cookies";
import { NextRequest } from "next/server";

export const proxyUser = async () => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens?.access_jwt) {
      return deleteAuthCookies();
    }

    const parsed = parseJwt(tokens.access_jwt);

    const response = await axios.get(
      `${process.env.NEXT_PUBLIC_API}/users/${parsed?.user_uid}`,
      {
        headers: { Authorization: `Bearer ${tokens.access_jwt}` },
      }
    );

    return Response.json(response.data, { status: 200 });
  } catch (error) {
    return publicApiErrorResponse(error);
  }
};

export const proxyUserAddresses = async () => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens?.access_jwt) {
      return deleteAuthCookies([]);
    }

    const parsed = parseJwt(tokens.access_jwt);

    const response = await axios.get(
      `${process.env.NEXT_PUBLIC_API}/users/${parsed?.user_uid}/delivery_address`,
      {
        headers: { Authorization: `Bearer ${tokens.access_jwt}` },
      }
    );

    return Response.json(response.data, { status: 200 });
  } catch (error) {
    return publicApiErrorResponse(error);
  }
};

export const proxyAddDeliveryAddress = async (req: NextRequest) => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens?.access_jwt) {
      return deleteAuthCookies();
    }

    const parsed = parseJwt(tokens.access_jwt);

    const body = await req.json();

    const response = await axios.post(
      `${process.env.NEXT_PUBLIC_API}/users/${parsed?.user_uid}/delivery_address`,
      { ...body },
      {
        headers: { Authorization: `Bearer ${tokens.access_jwt}` },
      }
    );

    return Response.json(response.data, { status: 200 });
  } catch (error) {
    return publicApiErrorResponse(error);
  }
};

export const proxyDeleteDeliveryAddress = async (req: NextRequest) => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens?.access_jwt) {
      return deleteAuthCookies();
    }

    const parsed = parseJwt(tokens.access_jwt);

    const body = await req.json();

    const response = await axios.delete(
      `${process.env.NEXT_PUBLIC_API}/users/${parsed?.user_uid}/delivery_address/${body.addressUid}`,
      {
        headers: { Authorization: `Bearer ${tokens.access_jwt}` },
      }
    );

    return Response.json(response.data, { status: 200 });
  } catch (error) {
    return publicApiErrorResponse(error);
  }
};

export const proxyDeleteAccount = async () => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens?.access_jwt) {
      return deleteAuthCookies();
    }

    const parsed = parseJwt(tokens.access_jwt);

    await axios.delete(`${process.env.NEXT_PUBLIC_API}/auth/user`, {
      headers: { Authorization: `Bearer ${tokens.access_jwt}` },
      data: {
        uid: parsed?.user_uid,
      },
    });

    return deleteAuthCookies();
  } catch (error) {
    return publicApiErrorResponse(error);
  }
};
