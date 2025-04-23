package helpers

import (
	"os"
	"runtime"
)

func HomeDir() string {
	homeDir := ""
	if runtime.GOOS == "windows" {
		homeDir = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if homeDir == "" {
			homeDir = os.Getenv("USERPROFILE")
		}
	}
	homeDir = os.Getenv("HOME")
	return homeDir
}
