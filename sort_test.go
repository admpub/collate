package collate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortStrings(t *testing.T) {
	ss := []string{`a_1`, `a_11`, `a_2`}
	SortStrings(ss)
	assert.Equal(t, []string{`a_1`, `a_2`, `a_11`}, ss)
}
