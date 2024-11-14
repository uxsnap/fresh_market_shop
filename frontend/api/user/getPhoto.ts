import client from "../client";

export const getPhoto = () => {
  return client.get<{ src: string }>("/user/photo");
};

getPhoto.queryKey = "getPhoto";
