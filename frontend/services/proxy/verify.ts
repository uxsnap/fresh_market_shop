import { getAuthCookieTokensFromServer } from "./cookies";

export const proxyVerify = async () => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();

    if (!tokens) {
      return Response.json({ isValid: false });
    }

    return Response.json({ isValid: true });
  } catch (err) {
    return Response.json({ isValid: false });
  }
};
