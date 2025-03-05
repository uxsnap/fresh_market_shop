import { dayJs } from "@/utils";

export const CART_MAIN_HEIGHT = 432;

export const API_PROXY_BASE_URL = "/api/proxy";

export const COOKIE_AUTH_TOKENS_NAME = "auth_tokens";

export const SECONDS_TO_EXPIRE = 15 * 1000;
export const MINUTE_IN_SECONDS = 60;
export const ONE_HOUR_IN_SECONDS = 60 * MINUTE_IN_SECONDS;
export const DAY_IN_SECONDS = 24 * ONE_HOUR_IN_SECONDS;
export const ONE_YEAR_IN_SECONDS = 365 * DAY_IN_SECONDS;

export const jwtError = "jwt_auth_middleware";

export const DEFAULT_MAP_ZOOM = 15;

export const COOKING_TIME_BORDERS = {
  min: dayJs.duration({ minutes: 5 }).asMilliseconds() * 1000,
  max: dayJs.duration({ hours: 6 }).asMilliseconds() * 1000,
};
