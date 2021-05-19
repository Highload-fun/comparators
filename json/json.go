package json

import (
	"io"

	"github.com/Highload-fun/comparators"
)

type Json struct{}

var _ comparators.Comparator = Json{}

func (Json) Compare(expected, got io.Reader) error {
	panic("Implement me")
}
