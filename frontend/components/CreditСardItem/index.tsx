import { ComponentProps } from "react";
import { SelectableItem } from "../SelectableItem";
import { CreditCard } from "../icons/CreditCard";

export const CreditCardItem = (
  props: Omit<ComponentProps<typeof SelectableItem>, "Icon">
) => <SelectableItem {...props} Icon={CreditCard} />;
