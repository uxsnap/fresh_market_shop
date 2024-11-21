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

export const YmapsWrapper = () => {
  const [coordinates, setCoordinates] = useState<number[][]>([[]]);

  const mapRef = useRef<ymaps.Map>();
  const map = useMapStore((s) => s.map);
  const setMap = useMapStore((s) => s.setMap);

  const setSearchValue = useMapStore((s) => s.setSearchValue);

  const handleAddress = async (coords: number[]) => {
    const geo = await map?.geocode(coords);
    const geoObject = geo?.geoObjects.get(0) as ExtendedGeoObject;

    // @ts-ignore
    console.log(geoObject.getLocalities(), geoObject.getThoroughfare());

    setSearchValue(geoObject.getThoroughfare());
    setCoordinates(coords as unknown as number[][]);
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

  useEffect(() => {
    setTimeout(() => {
      mapRef.current?.setCenter(DEFAULT_COORDS.center, DEFAULT_COORDS.zoom, {
        duration: 15,
      });
    }, 2000);
  }, [map]);

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
