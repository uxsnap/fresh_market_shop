import { useState } from "react";
import { Text, Image, SimpleGrid, Stack, CloseButton } from "@mantine/core";
import { Dropzone, IMAGE_MIME_TYPE, FileWithPath } from "@mantine/dropzone";
import { BackendImg } from "@/types";
import { useMutation } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { isServerImgFile, showErrorNotification } from "@/utils";
import { deleteProductPhoto } from "@/api/products/deleteProductPhoto";

const MAX_FILES = 3;

type Props = {
  productUid: string;
  files: (FileWithPath | BackendImg)[];
  setFiles: (files: (FileWithPath | BackendImg)[]) => void;
};

export const ImgsUpload = ({ productUid, files, setFiles }: Props) => {
  const handleFiles = (newFiles: FileWithPath[]) => {
    setFiles([...newFiles, ...files].slice(0, MAX_FILES));
  };

  const { mutate } = useMutation({
    mutationFn: deleteProductPhoto,
    mutationKey: [deleteProductPhoto.queryKey],
    onError: (error: AxiosError<any>) => {
      showErrorNotification(error);
    },
  });

  const handleDelete = (ind: number) => {
    const file = files[ind];

    const newArr = [...files];
    newArr.splice(ind, 1);
    setFiles(newArr);

    if (isServerImgFile(file)) {
      mutate({ uid: productUid, photoUid: file.uid });
    }
  };

  const previews = files.map((file, ind) => {
    const imageUrl = "uid" in file ? file.path : URL.createObjectURL(file);

    return (
      <Stack
        key={imageUrl}
        justify="center"
        align="center"
        pos="relative"
        mih={50}
        miw={50}
      >
        <Image
          src={imageUrl}
          onLoad={() =>
            "uid" in file ? file.path : URL.revokeObjectURL(imageUrl)
          }
        />

        <CloseButton
          pos="absolute"
          right={4}
          top={4}
          p={0}
          onClick={() => handleDelete(ind)}
          c="accent.0"
          size="xs"
        />
      </Stack>
    );
  });

  return (
    <Stack gap={12}>
      <Dropzone accept={IMAGE_MIME_TYPE} onDrop={handleFiles}>
        <Text ta="center">Загрузить изображения</Text>
      </Dropzone>

      <SimpleGrid cols={{ base: 1, sm: 4 }}>{previews}</SimpleGrid>
    </Stack>
  );
};
