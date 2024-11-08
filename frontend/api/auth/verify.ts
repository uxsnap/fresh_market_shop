import client from "../client";

type VerifyResponse = {
  isValid: boolean;
};

export const verifyUser = (): Promise<VerifyResponse> => {
  return client.post("/verify").then(data => data.data);
};

verifyUser.queryKey = "verifyUser";