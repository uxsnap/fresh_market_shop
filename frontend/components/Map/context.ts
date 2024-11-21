import { createFormContext } from "@mantine/form";

type MapForm = {
  city: string;
  addressUid: string;
  apartment?: number;
  entrance?: number;
  floor?: number;
  code?: number;
};

export const [MapFormProvider, useMapFormContext, useMapForm] =
  createFormContext<MapForm>();