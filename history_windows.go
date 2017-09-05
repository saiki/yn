// +build windows

package main

import (
	"bytes"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os/exec"
	"strings"
)

// history
func history() ([]string, error) {
	// sjis
	out, err := exec.Command("doskey", "/history").Output()
	if err != nil {
		return nil, err
	}
	sjisDecoder := japanese.ShiftJIS.NewDecoder()
	out, err = ioutil.ReadAll(transform.NewReader(bytes.NewReader(out), sjisDecoder))
	if err != nil {
		return nil, err
	}
	res := strings.Split(string(out), "\r\n")
	return res, nil
}
