import IDevice from "../interfaces/IDevice";

export const ApiSendDevice = async (device: IDevice) => {
  try {
    const res = await fetch("/api/devices", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(device),
    });
    console.log(res);
  } catch (err) {
    throw err;
  }
};

export const ApiGetDevices = async () => {
  try {
    const res = await fetch("/api/devices", {
      method: "GET",
      headers: { "Content-Type": "application/json" },
    });
    if (res.ok) {
      const json = (await res.json()) as IDevice[];
      return json;
    }
  } catch (err) {
    throw err;
  }
};
