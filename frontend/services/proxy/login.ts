import axios from "axios";
import { publicApiErrorResponse } from "@/utils";
import { NextRequest } from "next/server";
import {
  parseJwt,
  parseResponseCookies,
  setAuthCookiesFromResponse,
} from "./cookies";

export const proxyLogin = async (req: NextRequest) => {
  try {
    const body = await req.json();

    const response = await axios.post(
      `${process.env.NEXT_PUBLIC_API}/auth/login`,
      body
    );

    const parsed = parseResponseCookies(response);

    return setAuthCookiesFromResponse(response.data, parsed);
  } catch (error) {
    return publicApiErrorResponse(error);
  }
};
