import React, { useEffect, useState } from "react";
import {
  MapContainer,
  Marker,
  Polyline,
  Popup,
  TileLayer,
} from "react-leaflet";

import { ApiGetPositions, ApiGetPositionsByImei } from "../../api/ApiMapPoints";
import { ApiGetDevices } from "../../api/ApiDevices";

import ICoordinate from "../../interfaces/ICoordinate";
import IServerResponse from "../../interfaces/IServerResponse";
import IDevice from "../../interfaces/IDevice";
import IPosition from '../../interfaces/IPosition'

import { iconPerson } from "../Icons/Icons";
import "../App/App.css";

const MapPlaceholder: React.FC = () => {
  return (
    <p>
      Map of Ivanovo.{" "}
      <noscript>You need to enable JavaScript to see this map.</noscript>
    </p>
  );
};

// let startPosition = [{ lat: 56.99, lng: 40.97 }, { lat: 56.99, lng: 40.97}];
let center = { lat: 56.99, lng: 40.97 }
const zoom = 14;

const Map: React.FC = () => {
  const [positions, setPositions] = useState<ICoordinate[]>([]);
  const [devices, setDevices] = useState<IDevice[]>([]);
  const [startPosition, setStartPosition] = useState<ICoordinate[]>([])

  useEffect(() => {
    const getPos = async () => {
      const data: IServerResponse[] | undefined = await ApiGetPositions();
      if (data !== undefined) {
        const result: ICoordinate[] = data.map((item) => {
          return { lat: item.latitude, lng: item.longitude };
        });
        // startPosition = result;
        setPositions(result);
      }
    };
    getPos();

    const getDevices = async () => {
      const data: IDevice[] | undefined = await ApiGetDevices();
      if (data !== undefined) {
        const result: IDevice[] = data;
        setDevices(result);

        // let asyncRes: (IPosition[] | undefined)[] = []

        // try {
          const asyncRes = await Promise.all(result.map((value: IDevice)=>{
            return ApiGetPositionsByImei(value.imei)
          }))
        // } catch {}

        let position:ICoordinate[] = []

        asyncRes.forEach(function(value) {
          if(value !== undefined && value !== null){
            let t = value[0]
            let x = { lat: t.latitude, lng: t.longitude }
            position.push(x);
            () => setStartPosition(startPosition => position)
          }
        })
      }
    };
    getDevices();
  }, []);

  useEffect(()=>{
    console.log('change start position', startPosition)
  }, [startPosition])

  return (
    <MapContainer
      center={center}
      zoom={zoom}
      placeholder={MapPlaceholder}
    >
      <TileLayer
        attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
      />
      {devices.map((item: IDevice, index: number) => {
        return (
          <Marker position={center} icon={iconPerson}>
            <Popup
              closeButton={false}
              autoClose={false}
              closeOnEscapeKey={false}
              closeOnClick={false}
            >
              {item.name}
            </Popup>
          </Marker>
        );
      })}
      <Polyline positions={positions} />
    </MapContainer>
  );
};

export default Map;
