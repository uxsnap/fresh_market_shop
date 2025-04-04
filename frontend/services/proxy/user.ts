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

export const proxyDeleteAccount = async (req: NextRequest) => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens?.access_jwt) {
      return deleteAuthCookies();
    }

    let userUid;

    const body = await req.json();
    const isOtherUser = !!Object.keys(body).length;

    if (isOtherUser) {
      userUid = body.uid;
    } else {
      const parsed = parseJwt(tokens.access_jwt);

      userUid = parsed?.user_uid;
    }

    await axios.delete(`${process.env.NEXT_PUBLIC_API}/auth/user`, {
      headers: { Authorization: `Bearer ${tokens.access_jwt}` },
      data: {
        uid: userUid,
      },
    });

    if (!isOtherUser) {
      return deleteAuthCookies();
    }

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

export const proxySupportTickets = async () => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens?.access_jwt) {
      return deleteAuthCookies();
    }

    const parsed = parseJwt(tokens.access_jwt);

    const response = await axios.get(
      `${process.env.NEXT_PUBLIC_API}/support/tickets`,
      {
        params: {
          user_uid: parsed?.user_uid,
        },
        headers: { Authorization: `Bearer ${tokens.access_jwt}` },
      }
    );

    return Response.json(response.data, { status: 200 });
  } catch (error) {
    return publicApiErrorResponse(error);
  }
};
