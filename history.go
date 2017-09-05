package main

import (
	"errors"
)

func filter(yn string, history []string) (string, error) {
	if yn != "" {
		for i := len(history)-1; i > 1; i-- {
			if history[i] == yn {
				return history[i-1], nil
			}
		}
	}
	for i := len(history)-1; i > 0; i-- {
		if len(history[i]) > 0 {
			return history[i], nil
		}
	}
	return "", errors.New("no command history")
}
