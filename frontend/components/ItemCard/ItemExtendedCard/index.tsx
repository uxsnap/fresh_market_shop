// "use client";

// import {
//   Card,
//   Image,
//   Text,
//   Button,
//   Container,
//   Stack,
//   useMatches,
//   Box,
// } from "@mantine/core";
// import { ProductItem } from "@/types";
// import { Carousel } from "@mantine/carousel";
// import { getFallbackImg } from "@/utils";
// import { memo } from "react";
// import { useCartStore } from "@/store";

// import styles from "./ItemCard.module.css";
// import { ArrowsMinimize } from "@/components/icons/ArrowsMinimize";

// type Props = {
//   item: ProductItem;
// };


// // TODO: Remove unwanted rerenders thorough memoization of the state
// const ItemCounter = ({ item }: { item: ProductItem }) => {
//   const { incCartItem, decCartItem, addCartItem, getCount } = useCartStore();

//   const count = getCount(item.id);

//   return (
//     <Container fluid p={0} m={0} mt={8}>
//       {count === 0 ? (
//         <Button w="100%" onClick={() => addCartItem(item)} variant="accent">
//           Добавить
//         </Button>
//       ) : (
//         <Counter
//           count={count}
//           onDecrement={() => decCartItem(item.id)}
//           onIncrement={() => incCartItem(item.id)}
//         />
//       )}
//     </Container>
//   );
// };

// export const ItemCardExtended = memo(({ item }: Props) => {
//   const type = useMatches({
//     base: "small",
//     md: "default",
//   });

//   const { maw, imgH, priceFz, priceLh, infoFz, infoLh, nameFz, nameLh } =
//     mapTypeToValues[type];

//   const { price, name, imgs = [], info } = item;

//   const fallbackSrc = getFallbackImg(name);

//   return (
//     <Card p={8} w={maw} radius="md" withBorder pos="relative">
//       <Card.Section>
//         <ArrowsMinimize
//           className={styles.icon}
//           fill="var(--mantine-color-accent-0)"
//         />

//         {!imgs.length ? (
//           <Image
//             style={{ userSelect: "none" }}
//             src={""}
//             fallbackSrc={fallbackSrc}
//             height={imgH}
//             alt="Norway"
//             w="100%"
//           />
//         ) : (
//           <Carousel withControls={false}>
//             {imgs.map((img) => (
//               <Carousel.Slide key={img}>
//                 <Image
//                   style={{ userSelect: "none" }}
//                   loading="lazy"
//                   src={img}
//                   height={imgH}
//                   alt="Norway"
//                   fit="contain"
//                   fallbackSrc={fallbackSrc}
//                   w="100%"
//                 />
//               </Carousel.Slide>
//             ))}
//           </Carousel>
//         )}
//       </Card.Section>

//       <Stack mt={8} gap={4}>
//         <Text lh={`${priceLh}px`} fw={700} fz={priceFz} c="accent.0">
//           {price} Руб.
//         </Text>
//         <Text
//           truncate="end"
//           lh={`${infoLh}px`}
//           fw={500}
//           fz={infoFz}
//           c="accent.2"
//         >
//           {info}
//         </Text>
//       </Stack>

//       <Text
//         truncate="end"
//         lh={`${nameLh}px`}
//         fw={500}
//         fz={nameFz}
//         mt={8}
//         c="accent.0"
//       >
//         {name}
//       </Text>

//       <ItemCounter item={item} />
//     </Card>
//   );
// });
