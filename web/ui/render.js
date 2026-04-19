import { Events } from "../core/events.js";

function JsonSyntaxHighlight(json) {
  json = json
    .replace(/&/g, "&amp;")
    .replace(/</g, "&lt;")
    .replace(/>/g, "&gt;");

  return json.replace(
    /("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?)/g,
    function (match) {
      let cls = "number";

      if (/^"/.test(match)) {
        cls = /:$/.test(match) ? "key" : "string";
      } else if (/true|false/.test(match)) {
        cls = "boolean";
      } else if (/null/.test(match)) {
        cls = "null";
      }

      return `<span class="${cls}">${match}</span>`;
    },
  );
}

export function renderEvent(ev, onSelect) {
  const div = document.createElement("div");
  div.className = "event";

  // div.innerHTML = `
  //       <div class="kind">${Events[ev.kind] || ev.kind}</div>
  //       <div>${JsonSyntaxHighlight(ev.payload)}</div>
  //   `;
  div.innerHTML = `
        <div class="kind">${Events[ev.kind] || ev.kind}</div>
        <div>${JSON.stringify(ev.payload)}</div>
    `;

  div.onclick = () => onSelect(ev);
  return div;
}

export function showInspector(ev) {
  const el = document.getElementById("inspector");
  el.innerHTML = JsonSyntaxHighlight(JSON.stringify(ev, null, 2));
}

export function setStatus(text) {
  document.getElementById("status").textContent = text;
}

export function setUser(name) {
  document.getElementById("user").textContent = name;
}

export function addEvent(ev, onSelect) {
  const el = document.getElementById("events");
  el.appendChild(renderEvent(ev, onSelect));
}
