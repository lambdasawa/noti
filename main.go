package main

import (
	"os"
	"strings"

	"github.com/0xAX/notificator"
)

var notify *notificator.Notificator

func main() {
	message := strings.Join(os.Args[1:], "")

	notify = notificator.New(notificator.Options{})
	if err := notify.Push("", message, "", notificator.UR_CRITICAL); err != nil {
		panic(err)
	}
}
