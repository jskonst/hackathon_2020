import IServerRespone from "../interfaces/IServerResponse";

const url = "/api/positions";

export const ApiGetPositions = async () => {
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
