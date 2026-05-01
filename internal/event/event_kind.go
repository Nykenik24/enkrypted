package event

import (
	"encoding/json"
	"fmt"
	"strings"
)

type EventKind struct {
	definee   string
	namespace string
	name      string
}

func NewEventKind(definee, ns, name string) *EventKind {
	return &EventKind{
		definee:   definee,
		namespace: ns,
		name:      name,
	}
}

func (ek *EventKind) SetDefinee(definee string) *EventKind {
	ek.definee = definee
	return ek
}

func (ek *EventKind) SetNamespace(ns string) *EventKind {
	ek.namespace = ns
	return ek
}

func (ek *EventKind) SetName(name string) *EventKind {
	ek.name = name
	return ek
}

func (ek *EventKind) String() string {
	return fmt.Sprintf("%s:%s:%s", ek.definee, ek.namespace, ek.name)
}

func (ek *EventKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(ek.String())
}

func (ek *EventKind) CompareString(kindStr string) bool {
	return ek.String() == kindStr
}

func KindFromJSON(b []byte) (*EventKind, error) {
	var kind EventKind
	if err := json.Unmarshal(b, &kind); err != nil {
		return nil, err
	}

	return &kind, nil
}

func KindFromString(s string) (*EventKind, error) {
	parts := strings.SplitN(s, ":", 3)
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid format: expected definee:ns:name")
	}

	return NewEventKind(parts[0], parts[1], parts[2]), nil
}

type EventKindGroup struct {
	kinds []*EventKind
}

func NewEventKindGroup(initialKinds []*EventKind) *EventKindGroup {
	return &EventKindGroup{
		kinds: initialKinds,
	}
}

func (evg *EventKindGroup) Lookup(match string) (bool, *EventKind) {
	for _, kind := range evg.kinds {
		if kind.String() == match {
			return true, kind
		}
	}

	return false, nil
}

func (evg *EventKindGroup) indexOf(match string) (int, error) {
	for i, kind := range evg.kinds {
		if kind.String() == match {
			return i, nil
		}
	}

	return 0, fmt.Errorf("couldn't find kind `%s` in EventKindGroup", match)
}

func (evg *EventKindGroup) Add(kindStr string) *EventKindGroup {
	if exists, _ := evg.Lookup(kindStr); !exists {
		kind, err := KindFromString(kindStr)
		if err != nil {
			return evg
		}
		evg.kinds = append(evg.kinds, kind)
	}

	return evg
}

func (evg *EventKindGroup) Remove(kind string) *EventKindGroup {
	if exists, _ := evg.Lookup(kind); exists {
		i, err := evg.indexOf(kind)
		if err != nil {
			return evg
		}
		evg.kinds[i] = evg.kinds[len(evg.kinds)-1]
	}

	return evg
}
