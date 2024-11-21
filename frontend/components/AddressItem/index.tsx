import { ComponentProps } from "react";
import { LocationCursor } from "../icons/LocationCursor";

import { SelectableItem } from "../SelectableItem";

type Props = Omit<ComponentProps<typeof SelectableItem>, "Icon">;

export const AddressItem = (props: Props) => (
  <SelectableItem
    {...props}
    Icon={() => (
      <LocationCursor
        fill={`var(--mantine-color-${props.active ? "primary" : "accent"}-0)`}
      />
    )}
  />
);
