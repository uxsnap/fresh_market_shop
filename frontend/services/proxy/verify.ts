import {
  deleteAuthCookies,
  getAuthCookieTokensFromServer,
  isAccessTokenAlmostExpired,
  parseJwt,
} from "./cookies";
import { refresh } from "./refresh";

export const proxyVerify = async () => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens) {
      return deleteAuthCookies({ isValid: false, isAdmin: false });
    }

    const parsedJwt = parseJwt(tokens.access_jwt);

    if (isAccessTokenAlmostExpired(tokens.access_jwt)) {
      const res = await refresh(tokens);

      if (res !== true) {
        return deleteAuthCookies({
          isValid: false,
          isAdmin: parsedJwt?.role === "admin",
        });
      }
    }

    return Response.json({
      isValid: true,
      isAdmin: parsedJwt?.role === "admin",
    });
  } catch (err) {
    return deleteAuthCookies({ isValid: false, isAdmin: false });
  }
};
