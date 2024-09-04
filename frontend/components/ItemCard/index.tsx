import {
  Card,
  Image,
  Text,
  Button,
  Group,
  Container,
} from "@mantine/core";
import { useCounter } from "@mantine/hooks";
import { Counter } from "../Counter";

type Props = {
  type?: "default" | "small";
};

const mapTypeToValues = {
  default: {
    maw: 200,
    imgH: 176,
    priceFz: 22,
    priceLh: 26,
    infoFz: 12,
    infoLh: 14,
    nameFz: 14,
    nameLh: 16,
  },
  small: {
    maw: 140,
    imgH: 100,
    priceFz: 18,
    priceLh: 18,
    infoFz: 8,
    infoLh: 8,
    nameFz: 12,
    nameLh: 14,
  },
};

export const ItemCard = ({ type = "default" }: Props) => {
  const { maw, imgH, priceFz, priceLh, infoFz, infoLh, nameFz, nameLh } =
    mapTypeToValues[type];

  const [count, handlers] = useCounter(0, { min: 0, max: 10 });

  return (
    <Card p={8} maw={maw} radius="md" withBorder>
      <Card.Section>
        <Image
          src="https://raw.githubusercontent.com/mantinedev/mantine/master/.demo/images/bg-8.png"
          height={imgH}
          alt="Norway"
        />
      </Card.Section>

      <Group mt={8} gap={4}>
        <Text lh={`${priceLh}px`} fw={700} fz={priceFz} c="accent.0">
          2100 Руб.
        </Text>
        <Text lh={`${infoLh}px`} fw={500} fz={infoFz} c="accent.2">
          20 грамм/100 ккал.
        </Text>
      </Group>

      <Text lh={`${nameLh}px`} fw={500} fz={nameFz} mt={8} c="accent.0">
        Название товара
      </Text>

      <Container fluid p={0} m={0} mt={8}>
        {count === 0 ? (
          <Button w="100%" onClick={handlers.increment} variant="accent">
            Добавить
          </Button>
        ) : (
          <Counter
            count={count}
            onDecrement={handlers.decrement}
            onIncrement={handlers.increment}
          />
        )}
      </Container>
    </Card>
  );
};
