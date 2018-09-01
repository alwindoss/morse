package main

import (
	"fmt"
	"strings"

	"github.com/alwindoss/morse"
)

func main() {
	h := morse.NewHacker()
	morseCode, err := h.Encode(strings.NewReader("Convert this to Morse"))
	if err != nil {
		return
	}
	// Morse Code is: -.-. --- ...- . .-. - / - .... .. ... / - --- / -- --- .-. ... .
	fmt.Printf("Morse Code is: %s\n", string(morseCode))
}
