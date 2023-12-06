package clientinfo

import (
	"strings"
)

type OS int

type ClientInfo struct {
	OS      OS
	Version string
}

const (
	OS_Unknown OS = iota
	OS_iOS
	OS_Android
)

func ToOS(s string) OS {
	switch strings.ToLower(s) {
	case "ios":
		return OS_iOS
	case "android":
		return OS_Android
	}
	return OS_Unknown
}
