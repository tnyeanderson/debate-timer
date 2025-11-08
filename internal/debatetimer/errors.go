package debatetimer

import (
	"fmt"
)

// ErrorUnsupportedSpeaker is the error returned when a speaker number is not
// supported by the application. Only integers 1-9 are supported.
type ErrorUnsupportedSpeaker struct {
	input string
}

func NewErrorUnsupportedSpeaker(input string) ErrorUnsupportedSpeaker {
	return ErrorUnsupportedSpeaker{input}
}

func (e ErrorUnsupportedSpeaker) Error() string {
	return fmt.Sprintf("unsupported speaker %s, only speaker numbers 1-9 are supported", e.input)
}

// ErrorAlreadySpeaking is the error returned when a timer is started on a
// speaker who already the current speaker.
type ErrorAlreadySpeaking struct {
	speakerNumber int
}

func NewErrorAlreadySpeaking(speakerNumber int) ErrorAlreadySpeaking {
	return ErrorAlreadySpeaking{speakerNumber}
}

func (e ErrorAlreadySpeaking) Error() string {
	return fmt.Sprintf("%v is already speaking", GetSpeakerNameDefault(e.speakerNumber))
}
