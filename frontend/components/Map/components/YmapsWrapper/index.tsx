import { YMaps, Map as YandexMap, Placemark } from "@pbe/react-yandex-maps";
import {
  DEFAULT_COORDS_BY_CITY,
  DEFAULT_ZOOM,
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
import { useMapForm, useMapFormContext } from "../../context";

export const YmapsWrapper = () => {
  const [curCity, setCurCity] = useState("_");
  const [coordinates, setCoordinates] = useState<number[][]>([[]]);

  const form = useMapFormContext();

  form.watch("city", ({ value }) => {
    setCurCity(value);
  });

  const mapRef = useRef<ymaps.Map>();
  const map = useMapStore((s) => s.map);
  const setMap = useMapStore((s) => s.setMap);

  const setSearchValue = useMapStore((s) => s.setSearchValue);
  const mapAddress = useMapStore((s) => s.mapAddress);

  const handleAddress = async (coords: number[]) => {
    const geo = await map?.geocode(coords);
    const geoObject = geo?.geoObjects.get(0) as ExtendedGeoObject;

    const { houseNumber, street } = getStreetInfoFromGeo(geoObject);

    setSearchValue(`${street} ${houseNumber}`);
    setCoordinates(coords as unknown as number[][]);
  };

  const handleCenterMove = (coords: number[]) => {
    mapRef.current?.setCenter(coords, DEFAULT_ZOOM, {
      duration: 80,
    });
  };

  useEffect(() => {
    if (!map || !curCity) {
      return;
    }

    const coords = DEFAULT_COORDS_BY_CITY[curCity];

    setCoordinates(coords as unknown as number[][]);
    handleCenterMove(coords);
  }, [map, curCity]);

  useEffect(() => {
    if (!mapAddress) {
      return;
    }

    handleCenterMove([mapAddress.latitude, mapAddress.longitude]);
  }, [mapAddress]);

  const handleCoords = (e: any) => {
    const coords = e.get("coords");

    handleAddress(coords);
    setCoordinates(coords);
  };

  console.log(curCity);

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
        state={{
          center: DEFAULT_COORDS_BY_CITY[curCity],
          zoom: DEFAULT_ZOOM,
        }}
        onClick={handleCoords}
        onLoad={(ymaps: YMapsApi) => setMap(ymaps)}
        options={MAP_OPTIONS}
        instanceRef={mapRef}
      >
        {coordinates?.length && (
          <Placemark geometry={coordinates} options={PLACEMARK_OPTIONS} />
        )}
      </YandexMap>
    </YMaps>
  );
};
