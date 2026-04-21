export const Events = {
  "enkr:comm:msg": "Message",
  "enkr:room:identify": "Identify",
};

export function registerEvent(kind, label) {
  Events[kind] = label;
}
