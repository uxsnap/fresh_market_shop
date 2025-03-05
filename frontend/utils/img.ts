import { AxiosResponse } from "axios";

const drawImageOnCanvas = (img: HTMLImageElement): Promise<Blob> => {
  return new Promise((res, rej) => {
    const canvas = document.createElement("canvas");
    canvas.width = img.naturalWidth;
    canvas.height = img.naturalHeight;

    canvas.getContext("2d")!.drawImage(img, 0, 0);

    canvas.toBlob((blob) => {
      if (!blob) {
        rej();
        return;
      }

      res(blob);
    }, "image/webp");
  });
};

const createWebpFile = (fileName: string, blob: Blob) => {
  return new File([blob], `${fileName}.webp`, { type: blob.type });
};

export const processImgFile = (file: File): Promise<File> => {
  return new Promise((res, rej) => {
    if (!file) {
      rej();
      return;
    }

    const img = new Image();

    img.onload = () =>
      drawImageOnCanvas(img)
        .then((blob) => createWebpFile(file.name, blob))
        .then(res)
        .catch(rej);

    img.src = URL.createObjectURL(file);
  });
};

export const getBase64Img = (res: AxiosResponse<any, any>) => {
  return `data:${res.headers["content-type"]};base64,${Buffer.from(res.data).toString("base64")}`;
};

export const isServerImgFile = (file: { path?: string }) => {
  return "uid" in file;
};

export const renameFile = (originalFile: File, newName: string) => {
  return new File([originalFile], newName, {
    type: originalFile.type,
    lastModified: originalFile.lastModified,
  });
};

export const urlToObject = async (image: string, name: string) => {
  const response = await fetch(image);

  const blob = await response.blob();
  const file = new File([blob], name, { type: blob.type });

  return file;
};
