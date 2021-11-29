package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/0xAX/notificator"
)

var notify *notificator.Notificator

func main() {
	message := strings.Join(os.Args[1:], "")

	if canCreateDesktopNotification() {
		notify = notificator.New(notificator.Options{})
		escapedMessage := strings.ReplaceAll(message, `"`, `\"`)
		if err := notify.Push("", escapedMessage, "", notificator.UR_CRITICAL); err != nil {
			panic(err)
		}
	}
}

func canCreateDesktopNotification() bool {
	if _, err := exec.LookPath("notify-send"); err == nil {
		return true
	}

	if _, err := exec.LookPath("osascript"); err == nil {
		return true
	}

	return false
}
