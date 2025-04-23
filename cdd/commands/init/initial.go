package init

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/herryg91/cdd/cdd/pkg/configFile"
	"github.com/herryg91/cdd/cdd/pkg/helpers"
	"github.com/spf13/cobra"
)

type InitCmd struct {
	*cobra.Command
}

func NewInitCmd() *InitCmd {
	c := &InitCmd{}
	c.Command = &cobra.Command{
		Use:   "init",
		Short: "Initiate configuration",
		Long:  "Initiate configuration",
	}
	c.Command.RunE = c.runCommand
	return c
}

func (c *InitCmd) runCommand(cmd *cobra.Command, args []string) error {
	cf := configFile.ConfigFile{
		Extensions: []string{
			"$HOME/.cdd/ext/cddapis/",
			"$HOME/.cdd/ext/googleapis/",
		},
	}
	current_data, _ := json.MarshalIndent(cf, "", "    ")
	homeDir := helpers.HomeDir()

	if _, err := os.Stat(homeDir + "/.cdd"); !os.IsNotExist(err) {
		os.Mkdir(homeDir+"/.cdd", os.ModePerm)
	}

	if _, err := os.Stat(homeDir + "/.cdd/config"); !errors.Is(err, os.ErrNotExist) {
		fmt.Println("Config file is exist")
	} else {
		os.WriteFile(homeDir+"/.cdd/config", current_data, 0644)
	}

	// Put EXT
	cddext, err := readFromUrl("https://raw.githubusercontent.com/herryg91/cdd/refs/heads/main/protoc-gen-cdd/ext/cddapis/cdd/api/cddext.pb.go")
	if err != nil {
		fmt.Println("Failed to download cddext.pb.go")
		return nil
	}
	cddextproto, err := readFromUrl("https://raw.githubusercontent.com/herryg91/cdd/refs/heads/main/protoc-gen-cdd/ext/cddapis/cdd/api/cddext.proto")
	if err != nil {
		fmt.Println("Failed to download cddext.proto")
		return nil
	}
	annotationsproto, err := readFromUrl("https://raw.githubusercontent.com/herryg91/cdd/refs/heads/main/protoc-gen-cdd/ext/googleapis/google/api/annotations.proto")
	if err != nil {
		fmt.Println("Failed to download annotations.proto")
		return nil
	}
	httpproto, err := readFromUrl("https://raw.githubusercontent.com/herryg91/cdd/refs/heads/main/protoc-gen-cdd/ext/googleapis/google/api/http.proto")
	if err != nil {
		fmt.Println("Failed to download http.proto")
		return nil
	}
	httpbodyproto, err := readFromUrl("https://raw.githubusercontent.com/herryg91/cdd/refs/heads/main/protoc-gen-cdd/ext/googleapis/google/api/httpbody.proto")
	if err != nil {
		fmt.Println("Failed to download httpbody.proto")
		return nil
	}

	os.MkdirAll(homeDir+"/.cdd/ext/cddapis/cdd/api", os.ModePerm)
	os.MkdirAll(homeDir+"/.cdd/ext/googleapis/google/api", os.ModePerm)

	writeToFile(homeDir+"/.cdd/ext/cddapis/cdd/api/cddext.pb.go", cddext)
	writeToFile(homeDir+"/.cdd/ext/cddapis/cdd/api/cddext.proto", cddextproto)
	writeToFile(homeDir+"/.cdd/ext/googleapis/google/api/annotations.proto", annotationsproto)
	writeToFile(homeDir+"/.cdd/ext/googleapis/google/api/http.proto", httpproto)
	writeToFile(homeDir+"/.cdd/ext/googleapis/google/api/httpbody.proto", httpbodyproto)

	return nil
}

func readFromUrl(u string) (string, error) {
	response, err := http.Get(u)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Error: HTTP status code %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
		return "", fmt.Errorf("Error reading response body: %s", err.Error())

	}

	return string(body), nil
}

func writeToFile(path, content string) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(content)
	return nil
}
