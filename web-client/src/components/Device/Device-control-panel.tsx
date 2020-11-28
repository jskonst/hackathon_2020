import React, { useState } from "react";
import { ApiSendDevice } from "../../api/ApiDevices";
import "./Device.css";

type FormProps = {
  submit: () => void;
};

const Form: React.FC<FormProps> = ({ submit }) => {
  const [name, setName] = useState("");
  const [IMEI, setIMEI] = useState("");

  const handleChangeName = (event: any) => {
    const a = event.target.value;
    setName(a);
  };

  const handleChangeIMEI = (event: any) => {
    const a = event.target.value;
    if (a > 0) {
      setIMEI(a);
    }
  };

  const submitData = (event: any) => {
    if (name === "" || IMEI === "") {
      return;
    }
    const device = { name: name, imei: IMEI };
    ApiSendDevice(device);

    submit();
    event.preventDefault();
  };

  return (
    <>
      <form onSubmit={submitData}>
        <label>
          Device name:{" "}
          <input
            type="text"
            value={name}
            placeholder="0"
            onChange={handleChangeName}
          />
        </label>
        <br />
        <label>
          IMEI:{" "}
          <input
            type="number"
            value={IMEI}
            placeholder="0"
            onChange={handleChangeIMEI}
          />
        </label>
        <input className="pointer" type="submit" value="Send" />
      </form>
    </>
  );
};

const DeviceControlPanel: React.FC = () => {
  const [open, setOpen] = useState(false);

  function cover() {
    setOpen(!open);
  }

  if (open === false) {
    return (
      <div className="Device" onClick={cover}>
        Add device
      </div>
    );
  } else {
    return (
      <div className="Device">
        <Form submit={cover} />
        <p className="pointer" onClick={cover}>
          Close
        </p>
      </div>
    );
  }
};

export default DeviceControlPanel;
