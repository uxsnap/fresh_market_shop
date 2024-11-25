export const MAP_OPTIONS: ymaps.IMapOptions = {
  suppressMapOpenBlock: true,
  suppressObsoleteBrowserNotifier: true,
  copyrightLogoVisible: false,
  copyrightProvidersVisible: false,
  copyrightUaVisible: false,
};

export const MAP_MODULES = ["geolocation", "geocode"];

export interface ExtendedGeoObject extends ymaps.IGeoObject {
  getAddressLine(): string;
  getThoroughfare(): string;
}
