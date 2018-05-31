package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// Version returns current git tag if exists
func Version(gitdir string) string {
	cmd := exec.Command("git", "describe", "--always", "--tags")
	cmd.Dir = gitdir

	tag, err := cmd.Output()
	version := strings.TrimSpace(string(tag))

	if err != nil {
		// when no git tag exists
		fmt.Printf("err = %+v\n", err)
	}

	return version
}
