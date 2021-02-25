package gocli

import (
	"os/exec"
	"strings"
)

func GetCurrentModule() (string, error) {
	terminalCmd := exec.Command(`go`, "list", "-m")
	result, err := terminalCmd.Output()
	resultStr := strings.Replace(string(result), "\n", "", -1)
	return resultStr, err
}
