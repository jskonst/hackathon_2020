import React, { useEffect, useState } from "react";
import {
  MapContainer,
  Marker,
  Polyline,
  Popup,
  TileLayer,
} from "react-leaflet";
import { ApiGetPositions } from "../../api/ApiMapPoints";
import ICoordinate from "../../interfaces/ICoordinate";
import IServerResponse from "../../interfaces/IServerResponse";
import { iconPerson } from '../Icons/Icons'
import "../App/App.css";

const MapPlaceholder: React.FC = () => {
  return (
    <p>
      Map of Ivanovo.{" "}
      <noscript>You need to enable JavaScript to see this map.</noscript>
    </p>
  );
};

const startPosition = { lat: 56.99, lng: 40.97 };
const zoom = 14;

const fakePositions = [
  { lat: 56.23, lng: 40.217 },
  { lat: 56.35, lng: 40.357 },
  { lat: 56.42, lng: 40.467 },
  { lat: 56.56, lng: 40.527 },
  { lat: 56.67, lng: 40.697 },
  { lat: 56.79, lng: 40.757 },
  { lat: 56.99, lng: 40.97 },
];

const Map: React.FC = () => {
  const [positions, setPositions] = useState<ICoordinate[]>([]);

  useEffect(() => {
    const getPos = async () => {
      let data: IServerResponse[] | undefined = await ApiGetPositions();
      if (data !== undefined) {
        let result: ICoordinate[] = data.map((item) => {
          return { lat: item.latitude, lng: item.longitude };
        });
        setPositions(result);
      }
    };
    getPos();
  }, []);
  return (
    <MapContainer
      center={startPosition}
      zoom={zoom}
      placeholder={MapPlaceholder}
    >
      <TileLayer
        attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
      />
      <Marker
        position={startPosition}
        icon={iconPerson}
      >
        <Popup
          closeButton={false}
          autoClose={false}
          closeOnEscapeKey={false}
          closeOnClick={false}
        >
          Name very long for test
        </Popup>
      </Marker>
      <Polyline positions={positions} />
    </MapContainer>
  );
};
export default Map;
