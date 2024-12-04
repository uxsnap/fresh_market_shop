import { useMapStore } from "@/store/map";
import { Placemark } from "@pbe/react-yandex-maps";
import { useEffect, useState } from "react";
import { useMapFormContext } from "../../context";
import { DEFAULT_COORDS_BY_CITY } from "../../constants";

const PLACEMARK_OPTIONS = {
  iconLayout: "default#image",
  iconImageHref:
    "data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjQiIGhlaWdodD0iMjQiIHZpZXdCb3g9IjAgMCAyNCAyNCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPHBhdGggZD0iTTEzLjExMDkgMjMuNEMxNS41MTU2IDIwLjM5MDYgMjEgMTMuMDk2OSAyMSA5QzIxIDQuMDMxMjUgMTYuOTY4OCAwIDEyIDBDNy4wMzEyNSAwIDMgNC4wMzEyNSAzIDlDMyAxMy4wOTY5IDguNDg0MzggMjAuMzkwNiAxMC44ODkxIDIzLjRDMTEuNDY1NiAyNC4xMTcyIDEyLjUzNDQgMjQuMTE3MiAxMy4xMTA5IDIzLjRaTTEyIDZDMTIuNzk1NiA2IDEzLjU1ODcgNi4zMTYwNyAxNC4xMjEzIDYuODc4NjhDMTQuNjgzOSA3LjQ0MTI5IDE1IDguMjA0MzUgMTUgOUMxNSA5Ljc5NTY1IDE0LjY4MzkgMTAuNTU4NyAxNC4xMjEzIDExLjEyMTNDMTMuNTU4NyAxMS42ODM5IDEyLjc5NTYgMTIgMTIgMTJDMTEuMjA0NCAxMiAxMC40NDEzIDExLjY4MzkgOS44Nzg2OCAxMS4xMjEzQzkuMzE2MDcgMTAuNTU4NyA5IDkuNzk1NjUgOSA5QzkgOC4yMDQzNSA5LjMxNjA3IDcuNDQxMjkgOS44Nzg2OCA2Ljg3ODY4QzEwLjQ0MTMgNi4zMTYwNyAxMS4yMDQ0IDYgMTIgNloiIGZpbGw9IiM0RjQ2M0QiLz4KPC9zdmc+Cg==",
  iconImageSize: [48, 60],
  iconColor: "#4F463D",
};

type Props = {
  onCoords: (coords: number[]) => void;
};

export const MapPlacemark = ({ onCoords }: Props) => {
  const form = useMapFormContext();

  const deliveryAddress = useMapStore((s) => s.deliveryAddress);
  const mapAddress = useMapStore((s) => s.mapAddress);
  const mapInstance = useMapStore((s) => s.mapInstance);
  const handleCenterMove = useMapStore((s) => s.handleCenterMove);

  const [coordinates, setCoordinates] = useState<number[][]>([[]]);

  const handleCoords = (coords: number[][]) => {
    setCoordinates(coords);
    handleCenterMove(coords as unknown as number[]);

    onCoords(coords as unknown as number[]);
  };

  const handleMapEvents = (event: any) => {
    const coords = event.get("coords");

    handleCoords(coords);
  };

  useEffect(() => {
    if (!mapInstance) {
      return;
    }

    mapInstance?.events.add("click", handleMapEvents);

    return () => {
      mapInstance?.events.remove("click", handleMapEvents);
    };
  }, [mapInstance]);

  useEffect(() => {
    if (!mapAddress) {
      return;
    }

    handleCoords([
      mapAddress.latitude,
      mapAddress.longitude,
    ] as unknown as number[][]);
  }, [mapAddress]);

  useEffect(() => {
    if (!deliveryAddress) {
      return;
    }

    handleCoords([
      deliveryAddress.latitude,
      deliveryAddress.longitude,
    ] as unknown as number[][]);
  }, [deliveryAddress]);

  form.watch("city", ({ value }) => {
    const coords = DEFAULT_COORDS_BY_CITY[value];

    handleCoords(coords as unknown as number[][]);
  });

  if (!coordinates?.length) {
    return null;
  }

  return <Placemark geometry={coordinates} options={PLACEMARK_OPTIONS} />;
};
