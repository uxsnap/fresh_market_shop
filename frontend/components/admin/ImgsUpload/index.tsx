import { useState } from "react";
import { Text, Image, SimpleGrid, Stack, CloseButton } from "@mantine/core";
import { Dropzone, IMAGE_MIME_TYPE, FileWithPath } from "@mantine/dropzone";

const MAX_FILES = 3;

type Props = {
  files: FileWithPath[];
  setFiles: (files: FileWithPath[]) => void;
};

export const ImgsUpload = ({ files, setFiles }: Props) => {
  const handleFiles = (newFiles: FileWithPath[]) => {
    setFiles([...newFiles, ...files].slice(0, MAX_FILES));
  };

  const handleDelete = (ind: number) => {
    files.splice(ind, 1);

    setFiles([...files]);
  };

  const previews = files.map((file, ind) => {
    const imageUrl = URL.createObjectURL(file);

    return (
      <Stack justify="center" align="center" pos="relative" mih={50} miw={50}>
        <Image
          key={imageUrl}
          src={imageUrl}
          onLoad={() => URL.revokeObjectURL(imageUrl)}
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
