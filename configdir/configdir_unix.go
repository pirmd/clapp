// +build !windows,!darwin

package configdir

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	// SystemWide points to system-wide configuration path
	// On Unix systems, it is either /etc or first XDG_CONFIG_DIRS
	SystemWide = "/etc"

	// PerUser points to per-user configuration location
	// On Unix systems, it is either $HOME/.config or XDG_CONFIG_HOME
	PerUser = filepath.Join(os.Getenv("HOME"), ".config")
)

func init() {
	if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
		PerUser = xdg
	}

	if os.Getenv("XDG_CONFIG_DIRS") != "" {
		SystemWide = strings.Split(os.Getenv("XDG_CONFIG_DIRS"), ":")[0]
	}
}
