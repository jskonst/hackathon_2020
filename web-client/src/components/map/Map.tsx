import React, { useEffect, useState } from "react";
import { MapContainer, Marker, Popup, TileLayer } from "react-leaflet";
import { ApiMapPositions } from "../../api/ApiMapPoints";
import ICoordinate from "../../interfaces/ICoordinate";
import "../App/App.css";

const MapPlaceholder: React.FC = () => {
  return (
    <p>
      Map of London.{" "}
      <noscript>You need to enable JavaScript to see this map.</noscript>
    </p>
  );
};

let position = { lat: 56.99, lng: 40.97 };
const zoom = 14;

const Map: React.FC = () => {
  const [coords, coodrsMass] = useState({ lat: 56.99, lng: 40.97 });

  useEffect(() => {
    const getPos = async () => {
      const url = "http://localhost:3000/positions";
      let data = await ApiMapPositions(url);
      if (data !== undefined) {
        let x: ICoordinate = { lat: data[0].latituve, lng: data[0].longituve };
        coodrsMass((coords) => x);
      }
    }

    getPos();
  }, []);
  return (
    <MapContainer center={position} zoom={zoom} placeholder={MapPlaceholder}>
      <TileLayer
        attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
      />
      <Marker position={position}>
        <Popup>
          A pretty CSS3 popup. <br /> Easily customizable.
        </Popup>
      </Marker>
      {/* <Polyline positions ={positions} />  */}
    </MapContainer>
  );
};
export default Map;
