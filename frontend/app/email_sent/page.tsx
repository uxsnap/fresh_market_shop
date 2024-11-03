import { BackToCatalog } from "@/components/BackToCatalog";
import { HugeIconText } from "@/components/HugeIconText";

import styles from "./emailSent.module.css";
import { MainBox } from "@/components/MainBox";

export default function EmailSent() {
  return (
    <MainBox className={styles.root}>
      <BackToCatalog empty />

      <HugeIconText center type="email">
        Подтверждение регистрации отправлено на почту
      </HugeIconText>
    </MainBox>
  );
}
