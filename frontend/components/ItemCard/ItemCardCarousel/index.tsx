import { getFallbackImg } from "@/utils";
import { Carousel } from "@mantine/carousel";
import { Image } from "@mantine/core";

import styles from "./ItemCardCarousel.module.css";

type Props = {
  name: string;
  imgs: string[];
  className: string;
};

export const ItemCardCarousel = ({ name, imgs, className }: Props) => {
  const fallbackSrc = getFallbackImg(name);

  if (!imgs.length) {
    return (
      <Image
        style={{ userSelect: "none" }}
        src={""}
        fallbackSrc={fallbackSrc}
        className={className}
        alt="Norway"
        w="100%"
      />
    );
  }

  return (
    <Carousel withControls={false}>
      {imgs.map((img) => (
        <Carousel.Slide key={img}>
          <Image
            style={{ userSelect: "none" }}
            loading="lazy"
            src={img}
            className={className}
            alt="Norway"
            fit="contain"
            fallbackSrc={fallbackSrc}
            w="100%"
          />
        </Carousel.Slide>
      ))}
    </Carousel>
  );
};
