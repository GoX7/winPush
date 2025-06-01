package example

import (
	"log"

	"github.com/GoX7/winpush/winpush"
)

func action() {
	notify := winpush.Notificator{
		AppID: "App Name",
		Actions: []winpush.Actions{
			{
				Content:   "OK",
				Arguments: "action=ok",
			},
			{
				Content:   "Open page",
				Arguments: "https://google.com", // open website
			},
		},
	}

	if err := notify.Push(); err != nil {
		log.Print(err)
	}
}
