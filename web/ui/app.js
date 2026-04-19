import { createWS } from "../core/ws.js";
import { state } from "../core/state.js";
import { addEvent, showInspector, setStatus, setUser } from "./render.js";
import { Events } from "../core/events.js";

function initDropdown() {
  const select = document.getElementById("kind");

  select.innerHTML = "";

  Object.entries(Events).forEach(([key, label]) => {
    const opt = document.createElement("option");
    opt.value = key;
    opt.textContent = label;
    select.appendChild(opt);
  });
}

window.onload = () => {
  const url = document.getElementById("url");
  const username = document.getElementById("username");

  initDropdown();

  document.getElementById("connectBtn").onclick = () => {
    state.url = url.value;
    state.username = username.value || "";

    setUser(state.username);

    state.ws = createWS(state.url, (ev) => {
      state.events.push(ev);
      addEvent(ev, showInspector);
    });

    state.ws.onopen = () => {
      setStatus("connected");

      state.ws.send(
        JSON.stringify({
          kind: "enkr:identify",
          payload: { username: state.username },
        }),
      );
    };

    document.getElementById("connect").style.display = "none";
    document.getElementById("app").classList.remove("hidden");
  };

  document.getElementById("send").onclick = () => {
    if (!state.ws || state.ws.readyState !== 1) return;

    const kind = document.getElementById("kind").value;
    const msg = document.getElementById("msg").value;

    if (!kind) {
      console.warn("No event kind selected");
      return;
    }

    if (!msg.trim()) return;

    state.ws.send(
      JSON.stringify({
        kind: kind,
        payload: {
          contents: msg,
          timestamp: new Date().toISOString(),
        },
      }),
    );

    document.getElementById("msg").value = "";
  };
};
