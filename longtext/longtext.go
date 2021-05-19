package longtext

import (
	"io"

	"github.com/Highload-fun/comparators"
)

type LongText struct{}

var _ comparators.Comparator = LongText{}

func (LongText) Compare(expected, got io.Reader) error {
	panic("Implement me")
}
