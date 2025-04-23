package editor

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type EditorApp string

const (
	Editor_Vi   EditorApp = "vi"
	Editor_Nano EditorApp = "nano"
)

func (def EditorApp) FromString(in string) EditorApp {
	if strings.ToLower(in) == strings.ToLower(string(Editor_Nano)) {
		return Editor_Nano
	}
	return def
}

func Open(editorApp EditorApp, tmpFileName string, initData []byte) ([]byte, error) {
	if tmpFileName == "" {
		tmpFileName = "dply_tmp_" + strconv.FormatInt(time.Now().Unix(), 10)
	}
	tmpFile, err := ioutil.TempFile("", "scaling_edit")
	if err != nil {
		return initData, err
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Write(initData)
	tmpFile.Close()

	// open editor via terminal cmd
	termCmd := exec.Command(string(editorApp), tmpFile.Name())
	termCmd.Stdin = os.Stdin
	termCmd.Stdout = os.Stdout
	termCmd.Stderr = os.Stderr
	if err := termCmd.Run(); err != nil {
		return initData, err
	}

	return os.ReadFile(tmpFile.Name())
}
