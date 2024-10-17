import { Category } from "@/types";
import client from "../client";

export const getCategories = () => {
  return client.get<Category[]>("/categories");
};

getCategories.queryKey = "getCategories";
