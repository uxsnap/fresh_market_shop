import client from "../client";

export const getCategories = () => {
  return client.get("/categories");
};
