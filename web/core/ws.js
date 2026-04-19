export function createWS(url, onEvent) {
  const ws = new WebSocket(url);

  ws.onmessage = (e) => {
    try {
      onEvent(JSON.parse(e.data));
    } catch {
      onEvent({ kind: "raw", payload: e.data });
    }
  };

  return ws;
}
