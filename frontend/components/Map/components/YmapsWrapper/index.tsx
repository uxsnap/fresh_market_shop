import { YMaps, Map as YandexMap, Placemark } from "@pbe/react-yandex-maps";
import {
  DEFAULT_COORDS,
  ExtendedGeoObject,
  MAP_MODULES,
  MAP_OPTIONS,
  PLACEMARK_OPTIONS,
} from "./constants";
import { useEffect, useRef, useState } from "react";
import { YMapsApi } from "@pbe/react-yandex-maps/typings/util/typing";
import { useMapStore } from "@/store/map";
import ymaps from "yandex-maps";
import { getStreetInfoFromGeo } from "@/utils";

export const YmapsWrapper = () => {
  const [coordinates, setCoordinates] = useState<number[][]>([[]]);

  const mapRef = useRef<ymaps.Map>();
  const map = useMapStore((s) => s.map);
  const setMap = useMapStore((s) => s.setMap);

  const setSearchValue = useMapStore((s) => s.setSearchValue);
  const mapActiveAddress = useMapStore((s) => s.mapActiveAddress);

  const handleAddress = async (coords: number[]) => {
    const geo = await map?.geocode(coords);
    const geoObject = geo?.geoObjects.get(0) as ExtendedGeoObject;

    const { houseNumber, street } = getStreetInfoFromGeo(geoObject);

    setSearchValue(`${street} ${houseNumber}`);
    setCoordinates(coords as unknown as number[][]);
  };

  useEffect(() => {
    if (!map) {
      return;
    }

    setCoordinates(DEFAULT_COORDS.center as unknown as number[][]);
  }, [map]);

  useEffect(() => {
    if (!mapActiveAddress) {
      return;
    }

    mapRef.current?.setCenter(
      [mapActiveAddress.latitude, mapActiveAddress.longitude],
      DEFAULT_COORDS.zoom,
      {
        duration: 80,
      }
    );
  }, [mapActiveAddress]);

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
        width="100%"
        height="100%"
        modules={MAP_MODULES}
        defaultState={DEFAULT_COORDS}
        onClick={handleCoords}
        onLoad={(ymaps: YMapsApi) => setMap(ymaps)}
        options={MAP_OPTIONS}
        instanceRef={mapRef}
      >
        {coordinates.length && (
          <Placemark geometry={coordinates} options={PLACEMARK_OPTIONS} />
        )}
      </YandexMap>
    </YMaps>
  );
};
