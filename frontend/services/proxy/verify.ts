import { deleteAuthCookies, getAuthCookieTokensFromServer } from "./cookies";

export const proxyVerify = async () => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens) {
      return deleteAuthCookies({ isValid: false });
    }

    return Response.json({ isValid: true });
  } catch (err) {
    return deleteAuthCookies({ isValid: false });
  }
};
