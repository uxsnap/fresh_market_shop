import localFont from 'next/font/local';

export const robotoFont = localFont({
  src: [
    {
      path: './Roboto-Bold.ttf',
      weight: '700',
      style: 'normal',
    },
    {
      path: './Roboto-Medium.ttf',
      weight: '500',
      style: 'normal',
    },
    {
      path: './Roboto-Regular.ttf',
      weight: '400',
      style: 'normal',
    },
  ],
});