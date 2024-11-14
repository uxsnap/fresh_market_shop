import { Modal, Title } from "@mantine/core";
import { YMaps, Map as YandexMap, Placemark } from "@pbe/react-yandex-maps";

import styles from "./Map.module.css";
import { useEffect, useState } from "react";
import {
  DEFAULT_COORDS,
  MAP_MODULES,
  MAP_OPTIONS,
  PLACEMARK_OPTIONS,
} from "./constants";
import { MapFields } from "./components/MapFields";

type Props = {
  opened?: boolean;
  onClose: () => void;
};

export const Map = ({ opened = false, onClose }: Props) => {
  const [coordinates, setCoordinates] = useState<number[][]>([[]]);
  const [maps, setMaps] = useState<typeof YandexMap | null>(null);
  const [address, setAddress] = useState("");

  const handleClose = () => {
    close();
    onClose();
  };

  const handleAddress = async (coord: number[]) => {
    // @ts-ignore
    const geo = await maps?.geocode(coord);

    // @ts-ignore
    const geo1 = await maps?.geocode("Пятилеток 8к3");

    // @ts-ignore
    geo1.geoObjects.each((geoObj) => {
      console.log(geoObj.getAddressLine());
    });

    setAddress(geo.geoObjects.get(0).getAddressLine());
  };

  const onLoad = (map: any) => {
    setMaps(map);
  };

  useEffect(() => {
    if (!maps) {
      return;
    }

    handleAddress(DEFAULT_COORDS.center);

    // @ts-ignore For some reason this is the only way to make it work
    setCoordinates(DEFAULT_COORDS.center);
  }, [maps]);

  const handleCoords = (e: any) => {
    const coords = e.get("coords");

    handleAddress(coords);
    setCoordinates(coords);
  };

  return (
    <Modal.Root opened={opened} onClose={handleClose}>
      <Modal.Overlay />

      <Modal.Content className={styles.content}>
        <Modal.Header className={styles.header}>
          <Title c="accent.0"> Укажите ваш адрес</Title>
          <Modal.CloseButton size="32px" c="accent.0" />
        </Modal.Header>

        <Modal.Body px={20} py={12}>
          <MapFields />

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
              onLoad={(ymaps: any) => onLoad(ymaps)}
              options={MAP_OPTIONS}
            >
              {coordinates.length && (
                <Placemark geometry={coordinates} options={PLACEMARK_OPTIONS} />
              )}
            </YandexMap>
          </YMaps>
        </Modal.Body>
      </Modal.Content>
    </Modal.Root>
  );
};
