import React, { useEffect, useState } from "react";
import IDevice from "../../interfaces/IDevice";
import { ApiGetDevices } from "../../api/ApiDevices";
import "./Device.css";

export const DevicesList: React.FC = () => {
  const [devices, setDevices] = useState<IDevice[]>([]);
  useEffect(() => {
    const getDevices = async () => {
      let data: IDevice[] | undefined = await ApiGetDevices();
      if (data !== undefined) {
        let result: IDevice[] = data;
        setDevices(result);
      }
    };
    getDevices();
  }, []);

  return (
    <>
      {devices.map((item: IDevice) => {
        return (
          <ul className="Devices-list">
            <li>ID:{item.id}</li>
            <li>Name:{item.name}</li>
            <li>Imei:{item.imei}</li>
          </ul>
        );
      })}
    </>
  );
};
