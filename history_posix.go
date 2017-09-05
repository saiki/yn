// +build !windows

package main

import (
	"log"
	"os/exec"
	"strings"
)

func history() ([]string, error) {
	out, err := exec.Command("history").Output()
	if err != nil {
		return nil, err
	}
	log.Print(string(out))
	return strings.Split(string(out), "\n"), nil
}

