"use client";

import { BackToCatalog } from "@/components/BackToCatalog";
import { HugeIconText } from "@/components/HugeIconText";

import styles from "./emailVerify.module.css";
import { MainBox } from "@/components/MainBox";
import { useMutation } from "@tanstack/react-query";
import { verifyEmail } from "@/api/auth/verifyEmail";
import { useEffect, useState } from "react";
import { LoadingOverlay } from "@mantine/core";
import { useParams } from "next/navigation";

export default function EmailVerify() {
  const [verified, setVerified] = useState<boolean | undefined>();

  const params = useParams();
  const { token } = params;

  const mutation = useMutation({
    mutationFn: verifyEmail,
    onSuccess: () => {
      setVerified(true);
    },
    onError: () => {
      setVerified(false);
    },
  });

  useEffect(() => {
    mutation.mutate(token + "");
  }, [token]);

  return (
    <MainBox className={styles.root}>
      <BackToCatalog empty />

      <LoadingOverlay
        visible={verified === undefined}
        zIndex={1}
        overlayProps={{ radius: "sm", blur: 2 }}
        loaderProps={{ color: "primary.0", type: "bars" }}
      />

      {verified !== undefined && (
        <HugeIconText center type={verified ? "ok" : "sad"}>
          {verified
            ? "Вы успешно активировали свой аккаунт!"
            : "Не удалось активировать аккаунт!"}
        </HugeIconText>
      )}
    </MainBox>
  );
}
