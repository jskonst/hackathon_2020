import IServerRespone from "../interfaces/IServerResponse";
import IPosition from '../interfaces/IPosition'

export const ApiGetPositions = async () => {
  const url = "/api/positions";
  try {
    const res = await fetch(url, {
      method: "GET",
      headers: { "Content-Type": "application/json" },
    });
    if (res.ok) {
      const json = (await res.json()) as IServerRespone[];
      return json;
    }
  } catch (err) {
    throw err;
  }
};

export const ApiGetPositionsByImei = async (imei: string) => {
  try {
    const res = await fetch('/api/positions/' + imei, {
      method: "GET",
    });
    if (res.ok) {
      const json = (await res.json()) as IPosition[];
      return json;
    }
  } catch (err) {
    throw err;
  }
}
