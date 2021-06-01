package longtext

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/Highload-fun/comparators"
)

type LongText struct{}

var _ comparators.Comparator = LongText{}

func (LongText) Compare(expected, got io.Reader) error {
	expectedR := bufio.NewScanner(expected)
	gotR := bufio.NewScanner(got)

	line := 0
	for {
		expectedOk := expectedR.Scan()
		gotOk := gotR.Scan()
		if !expectedOk {
			if expectedR.Err() == nil {
				if !gotOk && gotR.Err() == nil {
					return nil // Files are equal
				}
				return fmt.Errorf("extra data was found")
			}

			return fmt.Errorf("cannot read expected data: %w", expectedR.Err())
		}

		if !gotOk {
			if gotR.Err() == nil {
				return fmt.Errorf("line %d: expected\n%s\ngot EOF", line, strings.TrimSpace(expectedR.Text()))
			}
			return gotR.Err()
		}

		expectedText := strings.TrimSpace(expectedR.Text())
		gotText := strings.TrimSpace(gotR.Text())
		if expectedText != gotText {
			return fmt.Errorf("line %d: expected\n%s\ngot\n%s", line, expectedText, gotText)
		}

		line++
	}
}
