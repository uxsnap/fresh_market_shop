import {
  proxyDefault,
  proxyLogin,
  proxyLogout,
  proxyVerify,
} from "@/services/proxy";
import { proxyGetPhoto, proxyUpdatePhoto } from "@/services/proxy/photo";
import { proxyUpdateUser } from "@/services/proxy/updateUser";
import { proxyUser } from "@/services/proxy/user";
import { NextRequest } from "next/server";

export async function GET(req: NextRequest) {
  const url = req.url?.replace(
    process.env.NEXT_PUBLIC_API_PROXY_BASE_URL + "",
    ""
  );

  switch (url) {
    case "/user/photo":
      return proxyGetPhoto();
    case "/user":
      return proxyUser();
  }

  return proxyDefault(req);
}

export async function POST(req: NextRequest) {
  const url = req.url?.replace(
    process.env.NEXT_PUBLIC_API_PROXY_BASE_URL + "",
    ""
  );

  switch (url) {
    case "/login":
      return proxyLogin(req);
    case "/verify":
      return proxyVerify();
    case "/logout":
      return proxyLogout();
    case "/user/photo":
      return proxyUpdatePhoto(req);
    default:
      const body = await req.json();
      return proxyDefault(req, body);
  }
}

export async function PUT(req: NextRequest) {
  const url = req.url?.replace(
    process.env.NEXT_PUBLIC_API_PROXY_BASE_URL + "",
    ""
  );

  switch (url) {
    case "/user":
      return proxyUpdateUser(req);
    default:
      const body = await req.json();
      return proxyDefault(req, body);
  }
}
