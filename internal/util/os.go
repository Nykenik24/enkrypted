package util

import "runtime"

type OSType int

const (
	Windows OSType = iota
	Unix
)

func GetRuntimeOS() OSType {
	if runtime.GOOS == "windows" {
		return Windows
	} else {
		return Unix
	}
}
