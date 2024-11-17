import { createFormContext } from "@mantine/form";

type MapForm = {
  flat?: number;
  entrance?: number;
  floor?: number;
  code?: number;
  address: string;
};

export const [MapFormProvider, useMapFormContext, useMapForm] =
  createFormContext<MapForm>();
