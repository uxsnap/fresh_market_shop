import axios from "axios";

type VerifyResponse = {
  isValid: boolean;
};

export const verifyUser = (): Promise<VerifyResponse> => {
  return axios.post(`${process.env.NEXT_PUBLIC_API_PROXY_BASE_URL}/verify`);
};

verifyUser.queryKey = "verifyUser";
