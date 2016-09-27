package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type Settings struct {
	Windows []string
	Linux   []string
	Mac     []string
}

func loadSettings(fpath string) (Settings, error) {
	set := Settings{}

	bytes, errRead := ioutil.ReadFile(fpath)
	if errRead != nil {
		return set, errRead
	}

	if errJson := json.Unmarshal(bytes, &set); errJson != nil {
		return set, errJson
	}

	return set, nil
}

func getSettingsPath() string {
	exePath := os.Args[0]
	exeName := strings.TrimSuffix(filepath.Base(exePath), filepath.Ext(exePath))
	jsonName := exeName + ".json"

	return filepath.Join(filepath.Dir(exePath), jsonName)
}

func SetWallpaper(fpath string, set Settings) error {
	cmds := []string{}

	var wrapCmd [2]string

	switch runtime.GOOS {
	case "windows":
		cmds = set.Windows
		wrapCmd = [2]string{"cmd", "/C"}
	case "linux":
		cmds = set.Linux
		wrapCmd = [2]string{"sh", "-c"}
	case "darwin":
		cmds = set.Mac
		wrapCmd = [2]string{"sh", "-c"}
	default:
		return errors.New("Unsupported OS")
	}

	for _, cmdArg := range cmds {
		cmdArg = strings.Replace(cmdArg, "__WALL__", fpath, -1)
		exec.Command(wrapCmd[0], wrapCmd[1], cmdArg).Run()
	}

	return nil
}
