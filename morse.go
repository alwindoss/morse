package morse

import (
	"bytes"
	"io"
	"io/ioutil"
	"strings"
)

// Encoder is an interface that specifies function to Encode
type Encoder interface {
	Encode(io.Reader) ([]byte, error)
}

// Decoder is an interface that specifies function to Encode
type Decoder interface {
	Decode(io.Reader) ([]byte, error)
}

// Hacker is an interface that composes Encoder and Decoder
type Hacker interface {
	Encoder
	Decoder
}

type hacker struct {
}

func (h *hacker) Encode(r io.Reader) ([]byte, error) {
	d, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	data := string(d)
	data = strings.ToUpper(data)
	var encode bytes.Buffer
	splitData := strings.Split(data, " ")
	numOfWords := len(splitData)
	for j, val := range splitData {
		vLen := len(val)
		for i, c := range val {
			char := string(c)
			if v, ok := alphaNumToMorse[char]; ok {
				encode.WriteString(v)
				if (i + 1) != vLen {
					encode.WriteByte(' ')
				}
			}
		}
		if numOfWords != (j + 1) {
			encode.WriteString(" / ")
		}
	}
	return encode.Bytes(), nil
}

func (h *hacker) Decode(r io.Reader) ([]byte, error) {
	d, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	data := string(d)

	splitData := strings.Split(data, " ")
	var decode bytes.Buffer
	for _, val := range splitData {
		if val == "/" {
			decode.WriteByte(' ')
		}
		decode.WriteString(morseToAlphaNum[val])
	}

	return decode.Bytes(), nil
}

// NewHacker is a factory function that generates a Morse Hacker Client
func NewHacker() Hacker {
	return &hacker{}
}
