import { createFormContext } from "@mantine/form";

export type MapFormInnerFields = {
  apartment?: number;
  entrance?: number;
  floor?: number;
  code?: number;
};

export type MapForm = {
  city: string;
  addressUid: string;
};

export const [MapFormProvider, useMapFormContext, useMapForm] =
  createFormContext<MapForm>();
