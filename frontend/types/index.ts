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
  products: ProductWithPhotos[];
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
  productUid: string;
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

export type User = {
  uid: string;
  firstName: string;
  lastName: string;
  birthday: string;
  email: string;
  createdAt: string;
  updatedAt: string;
};

export type ErrorWrapper = {
  type: string;
  message: string;
};

export type City = {
  uid: string;
  name: string;
};

export type CreditCard = {
  expired: string;
  externalUid: string;
  number: string;
  uid: string;
  userUid: string;
};

export type Address = {
  houseNumber: string;
  latitude: number;
  longitude: number;
  cityName?: string;
};

export type MapAddress = Address & {
  uid: string;
  cityUid: string;
  street: string;
};

export type DeliveryAddress = Address & {
  uid: string;
  userUid: string;
  addressUid: string;
  streetName: string;
  floors: number;
  entrances: number;
  apartment: number;
  codes: number;
  createdAt: Date;
  updatedAt: Date;
};

export type OrderProduct = {
  orderUid: string;
  productUid: string;
  count: number;
  photos: Photo[];
};

export type OrderStatus = "new" | "paid" | "in_progress" | "done";

export type Order = {
  uid: string;
  userUid: string;
  num: number;
  sum: number;
  status: OrderStatus;
  createdAt: string;
  updatedAt: string;
};

export type OrderWithProducts = {
  order: Order;
  products: OrderProduct[];
};

export type Topic = {
  uid: string;
  name: string;
  description: string;
};
