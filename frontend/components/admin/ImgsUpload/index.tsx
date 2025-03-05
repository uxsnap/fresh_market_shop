import { Text, Image, SimpleGrid, Stack, CloseButton } from "@mantine/core";
import { Dropzone, IMAGE_MIME_TYPE, FileWithPath } from "@mantine/dropzone";
import { BackendImg } from "@/types";

type Props = {
  files: (FileWithPath | BackendImg)[];
  setFiles: (files: (FileWithPath | BackendImg)[]) => void;
  onDelete: (file: FileWithPath | BackendImg) => void;
  onDrop: (files: FileWithPath[]) => void;
};

export const ImgsUpload = ({ files, setFiles, onDelete, onDrop }: Props) => {
  const handleDelete = (ind: number) => {
    const file = files[ind];

    const newArr = [...files];
    newArr.splice(ind, 1);
    setFiles(newArr);

    onDelete(file);
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
      <Dropzone accept={IMAGE_MIME_TYPE} onDrop={onDrop}>
        <Text ta="center">Загрузить изображения</Text>
      </Dropzone>

      <SimpleGrid cols={{ base: 1, sm: 4 }}>{previews}</SimpleGrid>
    </Stack>
  );
};
