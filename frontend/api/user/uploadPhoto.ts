import client from "../client";

export const uploadPhoto = (file: File) => {
  const formData = new FormData();

  formData.append("photo", file);

  return client.post("/user/photo", formData);
};

uploadPhoto.queryKey = "uploadPhoto";
