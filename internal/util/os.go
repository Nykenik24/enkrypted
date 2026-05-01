package util

import "runtime"

type OSType int

const (
	Windows OSType = iota
	Linux
	MacOS
	OpenBSD
	FreeBSD
	UnknownOS
)

var osToType = map[string]OSType{
	"windows": Windows,
	"darwin":  MacOS,
	"linux":   Linux,
	"openbsd": OpenBSD,
	"freebsd": FreeBSD,
}

var typeToString = (func() map[OSType]string {
	m := make(map[OSType]string)
	for k, v := range osToType {
		m[v] = k
	}
	return m
})()

func GetRuntimeOS() (ostype OSType) {
	var exists bool
	if ostype, exists = osToType[runtime.GOOS]; !exists {
		return UnknownOS
	}
	return ostype
}

func (t OSType) Unix() bool {
	return t == Linux || t == MacOS
}

func (t OSType) String() string {
	if t == UnknownOS {
		return "unknown"
	}
	return typeToString[t]
}
