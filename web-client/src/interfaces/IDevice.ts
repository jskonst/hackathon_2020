import ICoordinate from "./ICoordinate";
export default interface IDevice {
  id?: string;
  name: string;
  imei: string;
  positions?: ICoordinate[];
}
