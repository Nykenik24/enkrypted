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

function rfc3339(d) {
  function pad(n) {
    return n < 10 ? "0" + n : n;
  }

  function timezoneOffset(offset) {
    var sign;
    if (offset === 0) {
      return "Z";
    }
    sign = offset > 0 ? "-" : "+";
    offset = Math.abs(offset);
    return sign + pad(Math.floor(offset / 60)) + ":" + pad(offset % 60);
  }

  return (
    d.getFullYear() +
    "-" +
    pad(d.getMonth() + 1) +
    "-" +
    pad(d.getDate()) +
    "T" +
    pad(d.getHours()) +
    ":" +
    pad(d.getMinutes()) +
    ":" +
    pad(d.getSeconds()) +
    timezoneOffset(d.getTimezoneOffset())
  );
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
          kind: "enkr:room:identify",
          data: { username: state.username },
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
          timestamp: rfc3339(),
        },
      }),
    );

    document.getElementById("msg").value = "";
  };
};
