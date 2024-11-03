import client from "../client";

export const verifyEmail = (token: string) => {
  return client.post(`/auth/verify/email/${token}`);
};

verifyEmail.queryKey = "verifyEmail";
