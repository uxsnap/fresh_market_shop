import {
  deleteAuthCookies,
  getAuthCookieTokensFromServer,
  isAccessTokenAlmostExpired,
} from "./cookies";
import { refresh } from "./refresh";

export const proxyVerify = async () => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens) {
      return deleteAuthCookies({ isValid: false });
    }

    if (isAccessTokenAlmostExpired(tokens.access_jwt)) {
      const res = await refresh(tokens);

      if (res !== true) {
        return deleteAuthCookies({ isValid: false });
      }
    }

    return Response.json({ isValid: true });
  } catch (err) {
    return deleteAuthCookies({ isValid: false });
  }
};
