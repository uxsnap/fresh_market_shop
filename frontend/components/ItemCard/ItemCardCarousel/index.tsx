import { getFallbackImg } from "@/utils";
import { Carousel } from "@mantine/carousel";
import { Image } from "@mantine/core";

import styles from "./ItemCardCarousel.module.css";

type Props = {
  name: string;
  imgs: string[];
};

export const ItemCardCarousel = ({ name, imgs }: Props) => {
  const fallbackSrc = getFallbackImg(name);

  if (!imgs.length) {
    return (
      <Image
        style={{ userSelect: "none" }}
        src={""}
        fallbackSrc={fallbackSrc}
        className={styles.img}
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
            className={styles.img}
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
