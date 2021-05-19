package comparators

import "io"

type Comparator interface {
	Compare(expected, got io.Reader) error
}
