// Package omarchythemehookgo is a golang replacement for omarchy-theme-hook
//
// Honestly? A personal attempt at trying to fix broken themeing after Omarchy 3 -> 4
// Will stick to Go standard lib as much as possible
package main

import (
	"os/user"
	"path"
)

func getOmarchyCurrentThemeDirPath() string {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	return path.Join(currentUser.HomeDir, ".config", "omarchy", "current", "theme", "colors.toml")
}

func main() {
}
