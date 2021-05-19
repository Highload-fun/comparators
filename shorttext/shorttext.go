package shorttext

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/Highload-fun/comparators"
)

type ShortText struct{}

var _ comparators.Comparator = ShortText{}

func (ShortText) Compare(expected, got io.Reader) error {
	expectedData, err := ioutil.ReadAll(expected)
	if err != nil {
		return err
	}

	gotData, err := ioutil.ReadAll(got)
	if err != nil {
		return err
	}

	if !bytes.Equal(bytes.TrimSpace(expectedData), bytes.TrimSpace(gotData)) {
		return fmt.Errorf(`expected "%s", got "%s"`, string(expectedData), string(gotData))
	}

	return nil
}
