import { YMaps, Map as YandexMap, Placemark } from "@pbe/react-yandex-maps";
import {
  DEFAULT_COORDS,
  ExtendedGeoObject,
  MAP_MODULES,
  MAP_OPTIONS,
  PLACEMARK_OPTIONS,
} from "./constants";
import { useEffect, useState } from "react";
import { useMapFormContext } from "../../context";
import { YMapsApi } from "@pbe/react-yandex-maps/typings/util/typing";
import { useMapStore } from "@/store/map";

export const YmapsWrapper = () => {
  const [coordinates, setCoordinates] = useState<number[][]>([[]]);
  const setMap = useMapStore((s) => s.setMap);
  const map = useMapStore((s) => s.map);

  const form = useMapFormContext();

  const handleAddress = async (coords: number[]) => {
    const geo = await map?.geocode(coords);

    form.setFieldValue(
      "address",
      (geo?.geoObjects.get(0) as ExtendedGeoObject).getAddressLine()
    );
  };

  useEffect(() => {
    if (!map) {
      return;
    }

    handleAddress(DEFAULT_COORDS.center);
    setCoordinates(DEFAULT_COORDS.center as unknown as number[][]);
  }, [map]);

  const handleCoords = (e: any) => {
    const coords = e.get("coords");

    handleAddress(coords);
    setCoordinates(coords);
  };

  return (
    <YMaps
      query={{
        apikey: process.env.NEXT_PUBLIC_YMAP_API,
      }}
    >
      <YandexMap
        modules={MAP_MODULES}
        width="100%"
        height={450}
        defaultState={DEFAULT_COORDS}
        onClick={handleCoords}
        onLoad={(ymaps: YMapsApi) => setMap(ymaps)}
        options={MAP_OPTIONS}
      >
        {coordinates.length && (
          <Placemark geometry={coordinates} options={PLACEMARK_OPTIONS} />
        )}
      </YandexMap>
    </YMaps>
  );
};
