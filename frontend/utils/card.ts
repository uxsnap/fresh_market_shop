export const handleCardNumber = (number: string) => {
  let replaced = number.trim().replace(/[^0-9 ]/g, "");

  if ([4, 9, 14].includes(replaced.length)) {
    replaced += " ";
  }

  return replaced;
};

export const handleExpired = (expired: string, prevExpired: string) => {
  let replaced = expired.trim().replace(/[^0-9/]/g, "");

  if (replaced.length === 2 && prevExpired.at(-1) !== "/") {
    replaced += "/";
  }

  return replaced;
};
