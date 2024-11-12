import axios from "axios";
import {
  deleteAuthCookies,
  getAuthCookieTokensFromServer,
  parseJwt,
} from "./cookies";

export const proxyLogout = async () => {
  try {
    const { tokens } = await getAuthCookieTokensFromServer();
    const parsed = parseJwt(tokens?.access_jwt);

    const body = { uid: parsed?.user_uid };

    await axios.post(`${process.env.NEXT_PUBLIC_API}/auth/logout`, body, {
      headers: { Authorization: `Bearer ${tokens?.access_jwt}` },
    });
  } finally {
    return deleteAuthCookies();
  }
};
