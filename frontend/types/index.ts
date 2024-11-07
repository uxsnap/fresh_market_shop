export type IconProps = React.ComponentPropsWithoutRef<"svg"> & {
  size?: number;
  fill?: string;
};

export type Category = {
  createdAt: string;
  description: string;
  name: string;
  uid: string;
  updatedAt: string;
};

export type ProductsWithCategories = {
  products: Product[];
  categories: Category[];
};

export type Product = {
  uid: string;
  categoryUid: string;
  name: string;
  description: string;
  ccal: number;
  price: number;
  createdAt: string;
  updatedAt: string;
  weight: number;
};

export type Photo = {
  uid: string;
  path: string;
};

export type ProductWithPhotos = {
  product: Product;
  photos?: Photo[];
};

export type ProductItem = {
  id: string;
  price: number;
  name: string;
  imgs: string[];
  weight: number;
  ccal: number;
  description: string;
};

export type CartItem = {
  product: ProductItem;
  count: number;
};

export type MakeOrderItem = {
  product_uid: string;
  count: number;
};

export type AuthType = "login" | "reg" | "forgotPass" | "passRet";

export type RecipeStep = {
  recipeUid: string;
  step: number;
  description: string;
};

export type Recipe = {
  uid: string;
  name: string;
  createdAt: string;
  updatedAt: string;
  cookingTime: number;
  ccal: number;
};

export type NextServerResult = {
  success: boolean;
  tokens?: {
    access_jwt?: string;
    refresh_jwt?: string;
  };
  error?: Error | null;
};

export type JwtData = {
  email: string;
  expired_at: string;
  permissions: string;
  role: string;
  session_uid: string;
  user_uid: string;
};
