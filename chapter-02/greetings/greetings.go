package greetings

import (
    "fmt"
    "errors"
    "math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error){
    if name == "" {
        return "", errors.New("empty name")
    }
    message := fmt.Sprintf(randomFormat(), name)
    return message,nil
}

func randomFormat() string {
    // a slice of message formats

    formats := [] string {
        "Hi, %v. Welcome!",
        "Great to see you, %v.",
        "Hail, %v Well met!",
    }
    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]

}
