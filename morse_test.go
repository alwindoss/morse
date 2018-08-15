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
	expected := "-. .- -- . "
	if string(data) != expected {
		t.Errorf("Expected the value to be %s but was %s", expected, string(data))
		t.Fail()
	}
}

func TestEncodeWithSentence(t *testing.T) {
	h := morse.NewHacker()
	r := strings.NewReader("My Name is Alwin")
	data, err := h.Encode(r)
	if err != nil {
		t.Fail()
	}
	fmt.Println(string(data))
	t.Fail()
	// expected := "-. .- -- . "
	// if string(data) != expected {
	// 	t.Errorf("Expected the value to be %s but was %s", expected, string(data))
	// 	t.Fail()
	// }
}
