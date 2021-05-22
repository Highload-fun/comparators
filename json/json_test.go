package json

import (
	"os"
	"testing"
)

func TestJson_Compare(t *testing.T) {
	fExpected, err := os.Open("./expected.json")
	if err != nil {
		t.Fatal(err)
	}
	defer fExpected.Close()

	fGot, err := os.Open("./got_valid.json")
	if err != nil {
		t.Fatal(err)
	}
	defer fGot.Close()

	c := Json{}
	if err := c.Compare(fExpected, fGot); err != nil {
		t.Fatal(err)
	}
}

func TestJson_Compare_Err1(t *testing.T) {
	fExpected, err := os.Open("./expected.json")
	if err != nil {
		t.Fatal(err)
	}
	defer fExpected.Close()

	fGot, err := os.Open("./got_invalid1.json")
	if err != nil {
		t.Fatal(err)
	}
	defer fGot.Close()

	c := Json{}
	if err := c.Compare(fExpected, fGot); err == nil {
		t.Fatal("Must be error")
	}
}

func TestJson_Compare_Err2(t *testing.T) {
	fExpected, err := os.Open("./expected.json")
	if err != nil {
		t.Fatal(err)
	}
	defer fExpected.Close()

	fGot, err := os.Open("./got_invalid2.json")
	if err != nil {
		t.Fatal(err)
	}
	defer fGot.Close()

	c := Json{}
	if err := c.Compare(fExpected, fGot); err == nil {
		t.Fatal("Must be error")
	}
}

func TestJson_Compare_Err3(t *testing.T) {
	fExpected, err := os.Open("./expected.json")
	if err != nil {
		t.Fatal(err)
	}
	defer fExpected.Close()

	fGot, err := os.Open("./got_invalid3.json")
	if err != nil {
		t.Fatal(err)
	}
	defer fGot.Close()

	c := Json{}
	if err := c.Compare(fExpected, fGot); err == nil {
		t.Fatal("Must be error")
	}
}
