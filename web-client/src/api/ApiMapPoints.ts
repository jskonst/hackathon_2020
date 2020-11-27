import react from "react";
import IDataFormat from "../interfaces/ICoordinate";
import IServerRespone from "../interfaces/IServerResponse";

export const ApiMapPositions = async () => {
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
