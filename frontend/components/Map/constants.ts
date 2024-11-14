export const PLACEMARK_OPTIONS = {
  iconLayout: "default#image",
  iconImageHref:
    "data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjQiIGhlaWdodD0iMjQiIHZpZXdCb3g9IjAgMCAyNCAyNCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPHBhdGggZD0iTTEzLjExMDkgMjMuNEMxNS41MTU2IDIwLjM5MDYgMjEgMTMuMDk2OSAyMSA5QzIxIDQuMDMxMjUgMTYuOTY4OCAwIDEyIDBDNy4wMzEyNSAwIDMgNC4wMzEyNSAzIDlDMyAxMy4wOTY5IDguNDg0MzggMjAuMzkwNiAxMC44ODkxIDIzLjRDMTEuNDY1NiAyNC4xMTcyIDEyLjUzNDQgMjQuMTE3MiAxMy4xMTA5IDIzLjRaTTEyIDZDMTIuNzk1NiA2IDEzLjU1ODcgNi4zMTYwNyAxNC4xMjEzIDYuODc4NjhDMTQuNjgzOSA3LjQ0MTI5IDE1IDguMjA0MzUgMTUgOUMxNSA5Ljc5NTY1IDE0LjY4MzkgMTAuNTU4NyAxNC4xMjEzIDExLjEyMTNDMTMuNTU4NyAxMS42ODM5IDEyLjc5NTYgMTIgMTIgMTJDMTEuMjA0NCAxMiAxMC40NDEzIDExLjY4MzkgOS44Nzg2OCAxMS4xMjEzQzkuMzE2MDcgMTAuNTU4NyA5IDkuNzk1NjUgOSA5QzkgOC4yMDQzNSA5LjMxNjA3IDcuNDQxMjkgOS44Nzg2OCA2Ljg3ODY4QzEwLjQ0MTMgNi4zMTYwNyAxMS4yMDQ0IDYgMTIgNloiIGZpbGw9IiM0RjQ2M0QiLz4KPC9zdmc+Cg==",
  iconImageSize: [30, 42],
  iconColor: "#4F463D",
};

export const MAP_OPTIONS: ymaps.IMapOptions = {
  suppressMapOpenBlock: true,
  suppressObsoleteBrowserNotifier: true,
  copyrightLogoVisible: false,
  copyrightProvidersVisible: false,
  copyrightUaVisible: false,
};

export const DEFAULT_COORDS = {
  center: [59.934878873507, 30.318067359924303],
  zoom: 15,
};

export const MAP_MODULES = ["geolocation", "geocode"];
