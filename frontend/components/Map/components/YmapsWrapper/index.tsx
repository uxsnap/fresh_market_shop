import { YMaps, Map as YandexMap } from "@pbe/react-yandex-maps";
import { ExtendedGeoObject, MAP_MODULES, MAP_OPTIONS } from "./constants";
import { useState } from "react";
import { YMapsApi } from "@pbe/react-yandex-maps/typings/util/typing";
import { useMapStore } from "@/store/map";
import { getStreetInfoFromGeo } from "@/utils";
import { MapPlacemark } from "../MapPlacemark";
import { DEFAULT_MAP_ZOOM } from "@/constants";
import { DEFAULT_COORDS_BY_CITY } from "../../constants";

export const YmapsWrapper = () => {
  const [map, setMap] = useState<YMapsApi>();

  const curCity = useMapStore((s) => s.city);

  const setMapInstance = useMapStore((s) => s.setMapInstance);

  const setSearchValue = useMapStore((s) => s.setSearchValue);

  const handleAddress = async (coords: number[]) => {
    const geo = await map?.geocode(coords);
    const geoObject = geo?.geoObjects.get(0) as ExtendedGeoObject;

    const { houseNumber, street } = getStreetInfoFromGeo(geoObject);

    setSearchValue(`${street} ${houseNumber}`);
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
        state={{
          center: DEFAULT_COORDS_BY_CITY[curCity],
          zoom: DEFAULT_MAP_ZOOM,
        }}
        options={MAP_OPTIONS}
        onLoad={(ymaps) => setMap(ymaps)}
        instanceRef={(val) => setMapInstance(val)}
      >
        <MapPlacemark onCoords={handleAddress} />
      </YandexMap>
    </YMaps>
  );
};
