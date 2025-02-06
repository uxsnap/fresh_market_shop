import {
  proxyDefault,
  proxyLogin,
  proxyLogout,
  proxyVerify,
} from "@/services/proxy";
import { proxyGetOrder, proxyOrderHistory } from "@/services/proxy/order";
import {
  proxyAddPaymentCard,
  proxyGetPaymentCardsByUser,
} from "@/services/proxy/paymentCards";
import { proxyGetPhoto, proxyUpdatePhoto } from "@/services/proxy/photo";
import {
  proxyDeleteProduct,
  proxyDeleteProductPhoto,
  proxyProductPhotos,
  proxyReviveProduct,
} from "@/services/proxy/products";
import {
  proxyDeleteRecipe,
  proxyDeleteRecipePhotos,
  proxyRecipePhotos,
} from "@/services/proxy/recipes";
import { proxyUpdateUser } from "@/services/proxy/updateUser";
import {
  proxyAddDeliveryAddress,
  proxyDeleteAccount,
  proxyDeleteDeliveryAddress,
  proxyUser,
  proxyUserAddresses,
} from "@/services/proxy/user";
import { NextRequest } from "next/server";

export async function GET(req: NextRequest) {
  const params = req.nextUrl.searchParams;
  const url = req.url
    ?.replace(process.env.NEXT_PUBLIC_API_PROXY_BASE_URL + "", "")
    .replace(req.nextUrl.search, "");

  switch (url) {
    case "/user/photo":
      return proxyGetPhoto();
    case "/user":
      return proxyUser();
    case "/payments/cards/by_user":
      return proxyGetPaymentCardsByUser();
    case "/user/addresses":
      return proxyUserAddresses();
    case "/orders/history":
      return proxyOrderHistory();
    case "/orders":
      return proxyGetOrder(req, params);
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
      return proxyAddDeliveryAddress(req);
    case "/user/addresses/delete":
      return proxyDeleteDeliveryAddress(req);
    case "/user/delete":
      return proxyDeleteAccount(req);
    case "/products/delete":
      return proxyDeleteProduct(req);
    case "/recipes/delete":
      return proxyDeleteRecipe(req);
    case "/products/revive":
      return proxyReviveProduct(req);
    case "/products/photos":
      return proxyProductPhotos(req);
    case "/recipes/photos":
      return proxyRecipePhotos(req);
    case "/products/photos/delete":
      return proxyDeleteProductPhoto(req);
    case "/recipes/photos/delete":
      return proxyDeleteRecipePhotos(req);
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
