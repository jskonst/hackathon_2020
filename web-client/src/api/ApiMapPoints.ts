import react from "react";
import IDataFormat from "../interfaces/ICoordinate";
import IServerRespone from "../interfaces/IServerResponse";

export const ApiMapPositions = async (url: string) => {
  try {
    const res = await fetch(url, {
      method: "GET",
    });
    if (res.ok) {
      const json = (await res.json()) as IServerRespone[];
      console.log(json);
      return json;
    }
  } catch (err) {
    throw err;
  }
};
