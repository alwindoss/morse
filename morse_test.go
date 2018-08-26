package morse_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alwindoss/morse"
)

func TestEncode(t *testing.T) {
	h := morse.NewHacker()
	r := strings.NewReader("Name")
	data, err := h.Encode(r)
	if err != nil {
		t.Fail()
	}
	expected := "-. .- -- ."
	if string(data) != expected {
		t.Fatalf("Expected the value to be `%s` but was `%s`", expected, string(data))
	}
}

func TestEncodeNotValidLetter(t *testing.T) {
	h := morse.NewHacker()
	r := strings.NewReader("Hi,世界")
	data, err := h.Encode(r)
	if err != nil {
		t.Fail()
	}
	expected := ".... .. --..-- "
	if string(data) != expected {
		t.Fatalf("Expected the value to be \n`%s` but was \n`%s`", expected, string(data))
	}
}

func TestEncodeWithSentence(t *testing.T) {
	h := morse.NewHacker()
	r := strings.NewReader("My Name is Alwin Doss")
	data, err := h.Encode(r)
	if err != nil {
		t.Fail()
	}
	fmt.Println(string(data))
	expected := "-- -.-- / -. .- -- . / .. ... / .- .-.. .-- .. -. / -.. --- ... ..."
	if string(data) != expected {
		t.Fatalf("Expected the value to be \n`%s` but was \n`%s`", expected, string(data))
	}
}

func TestDecode(t *testing.T) {
	h := morse.NewHacker()
	r := strings.NewReader("-- -.-- / -. .- -- . / .. ... / .- .-.. .-- .. -. / -.. --- ... ...")
	alphaNum, err := h.Decode(r)
	if err != nil {
		t.Fail()
	}
	expected := "MY NAME IS ALWIN DOSS"
	if string(alphaNum) != expected {
		t.Fatalf("Expected the value to be \n`%s` but was \n`%s`", expected, string(alphaNum))
	}
}
