package bytes

import (
	"os"
	"testing"
)

func init() {
	if err := os.WriteFile("./expected.bin", []byte{0, 10, 20, 30, 40, 50, 60, 70, 80}, 0644); err != nil {
		panic(err)
	}
	if err := os.WriteFile("./got_valid.bin", []byte{0, 10, 20, 30, 40, 50, 60, 70, 80}, 0644); err != nil {
		panic(err)
	}
	if err := os.WriteFile("./got_invalid1.bin", []byte{0, 10, 20, 5, 40, 50, 60, 70, 80}, 0644); err != nil {
		panic(err)
	}
	if err := os.WriteFile("./got_invalid2.bin", []byte{0, 10, 20, 30, 40, 50}, 0644); err != nil {
		panic(err)
	}
	if err := os.WriteFile("./got_invalid3.bin", []byte{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}, 0644); err != nil {
		panic(err)
	}
}

func TestBytes_Compare(t *testing.T) {
	fExpected, err := os.Open("./expected.bin")
	if err != nil {
		t.Fatal(err)
	}
	defer fExpected.Close()

	fGot, err := os.Open("./got_valid.bin")
	if err != nil {
		t.Fatal(err)
	}
	defer fGot.Close()

	c := Bytes{}
	if err := c.Compare(fExpected, fGot); err != nil {
		t.Fatal(err)
	}
}

func TestBytes_Compare_Err1(t *testing.T) {
	fExpected, err := os.Open("./expected.bin")
	if err != nil {
		t.Fatal(err)
	}
	defer fExpected.Close()

	fGot, err := os.Open("./got_invalid1.bin")
	if err != nil {
		t.Fatal(err)
	}
	defer fGot.Close()

	c := Bytes{}
	if err := c.Compare(fExpected, fGot); err == nil {
		t.Fatal("Must be error")
	} else if err.Error() != "pos 3: expected 30 got 5" {
		t.Fatalf("Invalid message: %s", err)
	}
}

func TestBytes_Compare_Err2(t *testing.T) {
	fExpected, err := os.Open("./expected.bin")
	if err != nil {
		t.Fatal(err)
	}
	defer fExpected.Close()

	fGot, err := os.Open("./got_invalid2.bin")
	if err != nil {
		t.Fatal(err)
	}
	defer fGot.Close()

	c := Bytes{}
	if err := c.Compare(fExpected, fGot); err == nil {
		t.Fatal("Must be error")
	} else if err.Error() != "cannot read data: EOF" {
		t.Fatalf("Invalid message: %s", err)
	}
}

func TestBytes_Compare_Err3(t *testing.T) {
	fExpected, err := os.Open("./expected.bin")
	if err != nil {
		t.Fatal(err)
	}
	defer fExpected.Close()

	fGot, err := os.Open("./got_invalid3.bin")
	if err != nil {
		t.Fatal(err)
	}
	defer fGot.Close()

	c := Bytes{}
	if err := c.Compare(fExpected, fGot); err == nil {
		t.Fatal("Must be error")
	} else if err.Error() != "extra data was found" {
		t.Fatalf("Invalid message: %s", err)
	}
}
