import { Button, Group, Modal, TextInput, Title } from "@mantine/core";
import { YMaps, Map as YandexMap, Placemark } from "@pbe/react-yandex-maps";
import { Glass } from "../icons/Glass";

import styles from "./Map.module.css";

export const Map = () => {
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
              defaultState={{ center: [55.75, 37.57], zoom: 9 }}
            >
              <Placemark defaultProperties={{
                
              }} defaultGeometry={[55.751574, 37.573856]} />
            </YandexMap>
          </YMaps>
        </Modal.Body>
      </Modal.Content>
    </Modal.Root>
  );
};
