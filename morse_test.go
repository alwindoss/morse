package morse_test

import (
	"strings"
	"testing"

	"github.com/alwindoss/morse"
)

func TestEncode(t *testing.T) {
	h := morse.NewHacker()
	r := strings.NewReader("Name")

	data, err := h.Encode(r)

	assertNoError(t, err)
	assertEqual(t, "-. .- -- .", string(data))
}

func TestEncodeWithSentence(t *testing.T) {
	h := morse.NewHacker()
	r := strings.NewReader("My Name is Alwin Doss")

	data, err := h.Encode(r)

	assertNoError(t, err)
	assertEqual(t, "-- -.-- / -. .- -- . / .. ... / .- .-.. .-- .. -. / -.. --- ... ...", string(data))
}

func TestDecode(t *testing.T) {
	h := morse.NewHacker()
	r := strings.NewReader("-- -.-- / -. .- -- . / .. ... / .- .-.. .-- .. -. / -.. --- ... ...")

	alphaNum, err := h.Decode(r)

	assertNoError(t, err)
	assertEqual(t, "MY NAME IS ALWIN DOSS", string(alphaNum))
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fail()
	}
}

func assertEqual(t *testing.T, expected string, actual string) {
	if actual != expected {
		t.Errorf("Expected the value to be %s but was %s", expected, actual)
		t.Fail()
	}
}
