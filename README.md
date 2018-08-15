# morse
[![Go Report Card](https://goreportcard.com/badge/github.com/alwindoss/morse)](https://goreportcard.com/report/github.com/alwindoss/morse)
[![CircleCI](https://circleci.com/gh/alwindoss/morse.svg?style=svg)](https://circleci.com/gh/alwindoss/morse)
[![codecov](https://codecov.io/gh/alwindoss/morse/branch/master/graph/badge.svg)](https://codecov.io/gh/alwindoss/morse)
[![Documentation](https://godoc.org/github.com/alwindoss/morse?status.svg)](http://godoc.org/github.com/alwindoss/morse)
[![License](https://img.shields.io/pypi/l/Django.svg)](https://github.com/alwindoss/morse/blob/master/LICENSE)

Morse Code Library in Go

## Download and Use
`go get -u -v github.com/alwindoss/morse`

or

`dep ensure -add github.com/alwindoss/morse`

# Sample Usage

```
package main

import (
	"fmt"
	"strings"

	"github.com/alwindoss/morse"
)

func main() {
	h := morse.NewHacker()
	morseCode, err := h.Encode(strings.NewReader("Covert this to Morse"))
	if err != nil {
		return
	}
	// Morse Code is: -.-. --- ...- . .-. - / - .... .. ... / - --- / -- --- .-. ... .
	fmt.Printf("Morse Code is: %s\n", string(morseCode))
}
```
