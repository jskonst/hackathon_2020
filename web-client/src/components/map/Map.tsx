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

let startPosition = { lat: 56.99, lng: 40.97 };
const zoom = 14;

const Map: React.FC = () => {
  const [positions, setPositions] = useState<ICoordinate[]>([]);

  useEffect(() => {
    const getPos = async () => {
      const data: IServerResponse[] | undefined = await ApiGetPositions();
      if (data !== undefined) {
        const result: ICoordinate[] = data.map((item) => {
          return { lat: item.latitude, lng: item.longitude };
        });
        startPosition = result[0];
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
      <Marker position={startPosition} icon={iconPerson}>
        <Popup
          closeButton={false}
          autoClose={false}
          closeOnEscapeKey={false}
          closeOnClick={false}
        >
          Cool kitty
        </Popup>
      </Marker>
      <Polyline positions={positions} />
    </MapContainer>
  );
};
export default Map;
