import React from "react";
import Map from "../map/Map";
import { DevicesList } from "../Device/Device-view-table";
import DeviceControlPanel from "../Device/Device-control-panel";
import "./App.css";

const App: React.FC = () => {
  return (
    <>
      <DevicesList />
      <Map />
      <DeviceControlPanel />
    </>
  );
};

export default App;
