import {
  proxyDefault,
  proxyLogin,
  proxyLogout,
  proxyVerify,
} from "@/services/proxy";
import {
  proxyAddPaymentCard,
  proxyGetPaymentCardsByUser,
} from "@/services/proxy/paymentCards";
import { proxyGetPhoto, proxyUpdatePhoto } from "@/services/proxy/photo";
import { proxyUpdateUser } from "@/services/proxy/updateUser";
import {
  proxyAddUserAddress,
  proxyUser,
  proxyUserAddresses,
} from "@/services/proxy/user";
import { NextRequest } from "next/server";

export async function DELETE(req: NextRequest) {
  return proxyDefault(req);
}

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
    case "/payments/cards/by_user":
      return proxyGetPaymentCardsByUser();
    case "/user/addresses":
      return proxyUserAddresses();
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
    case "/payments/cards":
      return proxyAddPaymentCard(req);
    case "/user/addresses":
      return proxyAddUserAddress(req);
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
