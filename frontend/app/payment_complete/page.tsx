import { BackToCatalog } from "@/components/BackToCatalog";
import { HugeIconText } from "@/components/HugeIconText";

import styles from "./paymentComplete.module.css";
import { MainBox } from "@/components/MainBox";

export default function PaymentComplete() {
  return (
    <MainBox className={styles.root}>
      <BackToCatalog empty />

      <HugeIconText center type="ok">
        Оплата прошла успешно!
      </HugeIconText>
    </MainBox>
  );
}
