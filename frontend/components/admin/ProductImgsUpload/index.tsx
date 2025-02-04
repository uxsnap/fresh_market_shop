import { useState } from "react";
import { Text, Image, SimpleGrid, Stack, CloseButton } from "@mantine/core";
import { Dropzone, IMAGE_MIME_TYPE, FileWithPath } from "@mantine/dropzone";
import { BackendImg } from "@/types";
import { useMutation } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { isServerImgFile, showErrorNotification } from "@/utils";
import { deleteProductPhoto } from "@/api/products/deleteProductPhoto";
import { ImgsUpload } from "../ImgsUpload";

const MAX_FILES = 3;

type Props = {
  productUid: string;
  files: (FileWithPath | BackendImg)[];
  setFiles: (files: (FileWithPath | BackendImg)[]) => void;
};

export const ProductImgsUpload = ({ productUid, files, setFiles }: Props) => {
  const handleFiles = (uploaded: FileWithPath[]) => {
    const newFiles = [...uploaded, ...files];

    for (let i = MAX_FILES; i < newFiles.length; i++) {
      const file = newFiles[i];

      if (isServerImgFile(file)) {
        mutate({ uid: productUid, photoUid: file.uid });
      }
    }

    setFiles(newFiles.slice(0, MAX_FILES));
  };

  const { mutate } = useMutation({
    mutationFn: deleteProductPhoto,
    mutationKey: [deleteProductPhoto.queryKey],
    onError: (error: AxiosError<any>) => {
      showErrorNotification(error);
    },
  });

  const handleDelete = (file: FileWithPath | BackendImg) => {
    if (isServerImgFile(file)) {
      mutate({ uid: productUid, photoUid: file.uid });
    }
  };

  return (
    <ImgsUpload
      files={files}
      setFiles={setFiles}
      onDelete={handleDelete}
      onDrop={handleFiles}
    />
  );
};
