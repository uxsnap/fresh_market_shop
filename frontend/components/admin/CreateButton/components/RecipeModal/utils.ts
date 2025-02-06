import { COOKING_TIME_BORDERS } from "@/constants";
import { convertTimeToDuration } from "@/utils";

export const validateCookingTime = (value: string) => {
  if (value.length < 5) {
    return "Время приготовления не должно быть пустым";
  }

  const time = convertTimeToDuration(value);

  if (time < COOKING_TIME_BORDERS.min || time > COOKING_TIME_BORDERS.max) {
    return "Время приготовления выходит за временные границы";
  }

  return null;
};
