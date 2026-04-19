export const Events = {
  "enkr:msg": "Message",
  "enkr:identify": "Identify",
};

export function registerEvent(kind, label) {
  Events[kind] = label;
}
