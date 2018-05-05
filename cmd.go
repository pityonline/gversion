package main

import (
	"os/exec"
	"strings"
)

// Cmd run git commands, returns stdout, and an error if happens
func Cmd(gitdir string, arg ...string) (stdout string, err error) {
	cmd := exec.Command("git", arg...)
	cmd.Dir = gitdir

	out, err := cmd.Output()
	outString := strings.TrimSpace(string(out))

	return outString, err
}
