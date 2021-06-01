package longtext

import (
	"os"
	"testing"
)

func TestLongText_Compare(t *testing.T) {
	fExpected, err := os.Open("./expected.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer fExpected.Close()

	fGot, err := os.Open("./got_valid.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer fGot.Close()

	c := LongText{}
	if err := c.Compare(fExpected, fGot); err != nil {
		t.Fatal(err)
	}
}

func TestLongText_Compare_Err1(t *testing.T) {
	fExpected, err := os.Open("./expected.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer fExpected.Close()

	fGot, err := os.Open("./got_invalid1.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer fGot.Close()

	c := LongText{}
	if err := c.Compare(fExpected, fGot); err == nil {
		t.Fatal("Must be error")
	} else if err.Error() != "line 2: expected\ntext 3\ngot\ntext err" {
		t.Fatal("Invalid message")
	}
}

func TestLongText_Compare_Err2(t *testing.T) {
	fExpected, err := os.Open("./expected.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer fExpected.Close()

	fGot, err := os.Open("./got_invalid2.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer fGot.Close()

	c := LongText{}
	if err := c.Compare(fExpected, fGot); err == nil {
		t.Fatal("Must be error")
	} else if err.Error() != "line 4: expected\ntext 5\ngot EOF" {
		t.Fatal("Invalid message")
	}
}

func TestLongText_Compare_Err3(t *testing.T) {
	fExpected, err := os.Open("./expected.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer fExpected.Close()

	fGot, err := os.Open("./got_invalid3.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer fGot.Close()

	c := LongText{}
	if err := c.Compare(fExpected, fGot); err == nil {
		t.Fatal("Must be error")
	} else if err.Error() != "extra data was found" {
		t.Fatal("Invalid message")
	}
}
