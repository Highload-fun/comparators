package json

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/google/go-cmp/cmp"

	"github.com/Highload-fun/comparators"
)

type Json struct{}

var _ comparators.Comparator = Json{}

func (Json) Compare(expected, got io.Reader) error {
	expectedR := json.NewDecoder(expected)
	expectedR.UseNumber()
	gotR := json.NewDecoder(got)
	gotR.UseNumber()

	for {
		var expectedData, gotData interface{}
		expectedErr := expectedR.Decode(&expectedData)
		gotErr := gotR.Decode(&gotData)
		if expectedErr != nil {
			if expectedErr == io.EOF {
				if gotErr == io.EOF {
					return nil // Files are equal
				}
				return fmt.Errorf("extra data was found")
			}

			return fmt.Errorf("cannot read expected data: %w", expectedErr)
		}

		if gotErr != nil {
			return gotErr
		}

		if diff := cmp.Diff(expectedData, gotData); diff != "" {
			return fmt.Errorf("diff: %s", diff)
		}
	}
}
