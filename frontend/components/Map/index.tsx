import { Button, Group, Modal, TextInput, Title } from "@mantine/core";
import { YMaps, Map as YandexMap, Placemark } from "@pbe/react-yandex-maps";
import { Glass } from "../icons/Glass";

import styles from "./Map.module.css";
import { useState } from "react";

export const Map = () => {
  const [coordinates, setCoordinates] = useState<number[][]>([[]]);

  const mapOptions = {
    iconLayout: "default#image",
    iconImageHref:
      "data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjQiIGhlaWdodD0iMjQiIHZpZXdCb3g9IjAgMCAyNCAyNCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPHBhdGggZD0iTTEzLjExMDkgMjMuNEMxNS41MTU2IDIwLjM5MDYgMjEgMTMuMDk2OSAyMSA5QzIxIDQuMDMxMjUgMTYuOTY4OCAwIDEyIDBDNy4wMzEyNSAwIDMgNC4wMzEyNSAzIDlDMyAxMy4wOTY5IDguNDg0MzggMjAuMzkwNiAxMC44ODkxIDIzLjRDMTEuNDY1NiAyNC4xMTcyIDEyLjUzNDQgMjQuMTE3MiAxMy4xMTA5IDIzLjRaTTEyIDZDMTIuNzk1NiA2IDEzLjU1ODcgNi4zMTYwNyAxNC4xMjEzIDYuODc4NjhDMTQuNjgzOSA3LjQ0MTI5IDE1IDguMjA0MzUgMTUgOUMxNSA5Ljc5NTY1IDE0LjY4MzkgMTAuNTU4NyAxNC4xMjEzIDExLjEyMTNDMTMuNTU4NyAxMS42ODM5IDEyLjc5NTYgMTIgMTIgMTJDMTEuMjA0NCAxMiAxMC40NDEzIDExLjY4MzkgOS44Nzg2OCAxMS4xMjEzQzkuMzE2MDcgMTAuNTU4NyA5IDkuNzk1NjUgOSA5QzkgOC4yMDQzNSA5LjMxNjA3IDcuNDQxMjkgOS44Nzg2OCA2Ljg3ODY4QzEwLjQ0MTMgNi4zMTYwNyAxMS4yMDQ0IDYgMTIgNloiIGZpbGw9IiM0RjQ2M0QiLz4KPC9zdmc+Cg==",
    iconImageSize: [30, 42],
    iconColor: "#4F463D",
  };

  return (
    <Modal.Root opened={true} onClose={close}>
      <Modal.Overlay />

      <Modal.Content miw={800}>
        <Modal.Header className={styles.header} px={20} py={12}>
          <Title c="accent.0"> Укажите ваш адрес</Title>
          <Modal.CloseButton size="32px" c="accent.0" />
        </Modal.Header>

        <Modal.Body px={20} py={12}>
          <Group grow gap={16}>
            <Group grow>
              <TextInput
                size="md"
                leftSection={<Glass size={16} />}
                placeholder="Поиск товаров"
              />
            </Group>

            <Button px={4} fz={12} maw={110} variant="accent">
              Добавить адрес
            </Button>

            <Button px={4} fz={12} maw={110} variant="accent-reverse">
              Выбрать адрес
            </Button>
          </Group>

          <Group mah={54} grow my={16} gap={20}>
            <TextInput
              lh={1}
              size="md"
              label="Квартира"
              placeholder="Введите квартиру"
            />

            <TextInput
              lh={1}
              size="md"
              label="Подъезд"
              placeholder="Введите подъезд"
            />

            <TextInput
              lh={1}
              size="md"
              label="Этаж"
              placeholder="Введите этаж"
            />

            <TextInput
              lh={1}
              size="md"
              label="Домофон"
              placeholder="Введите домофон"
            />
          </Group>

          <YMaps>
            <YandexMap
              width="100%"
              height={450}
              onClick={(e: any) => setCoordinates(e.get("coords"))}
              defaultState={{ center: [55.75, 37.57], zoom: 9 }}
            >
              {coordinates.length && (
                <Placemark geometry={coordinates} options={mapOptions} />
              )}
            </YandexMap>
          </YMaps>
        </Modal.Body>
      </Modal.Content>
    </Modal.Root>
  );
};
