package log

import (
	"log"

	"github.com/pkg/errors"
)

func OrigirnalError() error {
	return errors.New("error occurred")
}

func PassThroughError() error {
	err := OrigirnalError()

	return errors.Wrap(err, "in passthrouhError")
}

func FinalDestination() {
	err := PassThroughError()
	if err != nil {
		log.Printf("an error occurred: %s\n", err.Error())
		return
	}
}
