package utils

import (
	"bytes"
	"os/exec"
	"strings"
)

func IsRunning() bool {
	cmd := exec.Command("/bin/bash", "-c", "service troop-scout status")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		//fmt.Printf("systemctl status filebeat.service failed: %s", stderr.String())
		//	fmt.Print("down")
	}

	if strings.Contains(string(out.String()), "Active: active") {
		return true
	} else {
		return false
	}
}
