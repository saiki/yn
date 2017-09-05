package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {

	his, err := history()
	if err != nil {
		log.Fatal(err)
	}
	var yn string
	if len(os.Args) > 1 {
		yn = os.Args[1]
	}

	latest, err := filter(yn, his)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(latest)
	command, args := splitCommandAndArgs(latest)
	if command == "yn" {
		os.Exit(0)
	}
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		if e2, ok := err.(*exec.ExitError); ok {
			if s, ok := e2.Sys().(syscall.WaitStatus); ok {
				os.Exit(s.ExitStatus())
			} else {
				log.Fatal(err)
			}
		}
	}
	os.Exit(0)
}

func splitCommandAndArgs(cmd string) (string, []string) {
	v := strings.Split(cmd, " ")
	return v[0], v[1:]
}
