import { Carousel } from "@mantine/carousel";
import { Image } from "@mantine/core";

type Props = {
  name: string;
  imgs: string[];
  className: string;
};

export const ItemCardCarousel = ({ name, imgs, className }: Props) => (
  <Carousel withControls={false}>
    {imgs.map((img) => (
      <Carousel.Slide key={img}>
        <Image
          style={{ userSelect: "none" }}
          loading="lazy"
          src={img}
          className={className}
          alt={name}
          fit="contain"
          fallbackSrc={img}
          w="100%"
        />
      </Carousel.Slide>
    ))}
  </Carousel>
);
