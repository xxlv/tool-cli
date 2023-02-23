package main

import (
	"os/exec"
)

func execShell(cmd string, args ...string) (string, error) {
	r := exec.Command(cmd, args...)
	out, err := r.Output()

	return string(out), err
}

func runShell(cmd string, args ...string) (string, error) {
	r := exec.Command(cmd, args...)
	err := r.Run()
	return "", err
}
