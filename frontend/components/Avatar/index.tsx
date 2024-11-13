"use client";

import {
  FileButton,
  Avatar as MantineAvatar,
  LoadingOverlay,
} from "@mantine/core";

import { User } from "../icons/User";
import { processImgFile } from "@/utils";

import styles from "./Avatar.module.css";
import { useMutation, useQuery } from "@tanstack/react-query";
import { uploadPhoto } from "@/api/user/uploadPhoto";
import { useRef } from "react";
import { getPhoto } from "@/api/user/getPhoto";

type Props = {
  src?: string;
  size?: "small" | "default";
  upload?: boolean;
};

export const Avatar = ({ size = "default", upload = false }: Props) => {
  const resetRef = useRef<() => void>(null);
  const clearFile = () => resetRef.current?.();

  const { data, refetch, isFetching } = useQuery({
    queryFn: getPhoto,
    queryKey: [getPhoto.queryKey],
  });

  const { mutate, isPending } = useMutation({
    mutationFn: uploadPhoto,
    onSuccess: () => {
      refetch();
      clearFile();
    },
    onError: () => {
      clearFile();
    },
  });

  const handleUpload = async (file: File | null) => {
    if (!file) {
      return;
    }

    const newFile = await processImgFile(file);

    mutate(newFile);
  };

  return (
    <FileButton
      resetRef={resetRef}
      disabled={!upload || isFetching || isPending}
      onChange={handleUpload}
      accept="image/png,image/jpeg,image/jpg,image/webp"
    >
      {(props) => (
        <MantineAvatar
          className={styles.root}
          src={data?.data.src ?? ""}
          c="bg.1"
          size={size === "default" ? 250 : 38}
          {...props}
        >
          <LoadingOverlay
            visible={isFetching || isPending}
            zIndex={1}
            overlayProps={{ radius: "sm", blur: 2 }}
            loaderProps={{ color: "primary.0", type: "bars" }}
          />

          <User
            fill={`var(--mantine-color-accent-2`}
            size={size === "default" ? 100 : 19}
          />
        </MantineAvatar>
      )}
    </FileButton>
  );
};
