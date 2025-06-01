package example

import (
	"log"

	"github.com/GoX7/winpush/winpush"
)

func base() {
	notify := winpush.Notificator{
		AppID:               "App Name",
		Title:               "Title",
		Subtitle:            "Subtitle",
		Message:             "Message",
		Icon:                "path",               // path to image
		ActivationArguments: "https://google.com", //open website by clicking on notification
		Duration:            "short",              // execution time (short/long)
	}

	if err := notify.Push(); err != nil {
		log.Print(err)
	}
}
