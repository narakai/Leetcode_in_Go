package insert_sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	ast := assert.New(t)

	test1 := []int{1, 8, 5, 9, 3}
	test2 := []int{1, 3, 5, 8, 9}

	ast.Equal(test2, insert(test1))
}
