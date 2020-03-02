package helpers

import "log"

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
