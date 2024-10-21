import { BackToCatalog } from "@/components/BackToCatalog";
import { CartLeft } from "@/components/pages/cart/CartLeft";
import { Container } from "@mantine/core";

export default function CartPage() {
  return (
    <Container pt={22} m={0} maw={1454} mx="auto">
      <BackToCatalog />

      <CartLeft />
    </Container>
  );
}
