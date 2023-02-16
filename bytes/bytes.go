package bytes

import (
	"bufio"
	"errors"
	"fmt"
	"io"

	"github.com/Highload-fun/comparators"
)

type Bytes struct{}

var _ comparators.Comparator = Bytes{}

func (Bytes) Compare(expected, got io.Reader) error {
	expectedR := bufio.NewReader(expected)
	gotR := bufio.NewReader(got)

	pos := 0
	for {
		expectedByte, expectedErr := expectedR.ReadByte()
		gotByte, gotErr := gotR.ReadByte()

		if expectedErr != nil {
			if errors.Is(expectedErr, io.EOF) {
				if errors.Is(gotErr, io.EOF) {
					return nil
				}
				return fmt.Errorf("extra data was found")
			}
			return fmt.Errorf("cannot read expected data: %w", expectedErr)
		}

		if gotErr != nil {
			return fmt.Errorf("cannot read data: %w", gotErr)
		}

		if expectedByte != gotByte {
			return fmt.Errorf("pos %d: expected %d got %d", pos, expectedByte, gotByte)
		}

		pos++
	}
}
