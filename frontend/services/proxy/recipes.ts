import { NextRequest } from "next/server";
import { deleteAuthCookies, getAuthCookieTokensFromServer } from "./cookies";
import axios from "axios";
import { publicApiErrorResponse } from "@/utils";

export const proxyDeleteRecipe = async (req: NextRequest) => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens?.access_jwt) {
      return deleteAuthCookies();
    }

    const body = await req.json();

    await axios.delete(`${process.env.NEXT_PUBLIC_API}/recipes/${body.uid}`, {
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

export const proxyDeleteRecipePhotos = async (req: NextRequest) => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens?.access_jwt) {
      return deleteAuthCookies();
    }

    const body = await req.json();

    await axios.delete(
      `${process.env.NEXT_PUBLIC_API}/recipes/${body.uid}/photos`,
      {
        data: { photos: body.photos },
        headers: { Authorization: `Bearer ${tokens.access_jwt}` },
      }
    );

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

export const proxyDeleteRecipeStep = async (req: NextRequest) => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens?.access_jwt) {
      return deleteAuthCookies();
    }

    const body = await req.json();

    await axios.delete(
      `${process.env.NEXT_PUBLIC_API}/recipes/${body.uid}/steps/${body.step}`,
      {
        headers: { Authorization: `Bearer ${tokens.access_jwt}` },
      }
    );

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

export const proxyRecipePhotos = async (req: NextRequest) => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens?.access_jwt) {
      return deleteAuthCookies();
    }

    const formData = await req.formData();
    const uid = formData.get("uid");

    await axios.postForm(
      `${process.env.NEXT_PUBLIC_API}/recipes/${uid}/photos`,
      formData,
      {
        headers: { Authorization: `Bearer ${tokens.access_jwt}` },
      }
    );

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

export const proxyAddRecipeSteps = async (req: NextRequest) => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens?.access_jwt) {
      return deleteAuthCookies();
    }

    const body = await req.json();

    await axios.post(
      `${process.env.NEXT_PUBLIC_API}/recipes/${body.uid}/steps`,
      { steps: body.steps },
      {
        headers: { Authorization: `Bearer ${tokens.access_jwt}` },
      }
    );

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
