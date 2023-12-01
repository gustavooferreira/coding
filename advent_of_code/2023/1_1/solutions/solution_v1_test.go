package solutions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent_of_code/2023/1_1/solutions"
)

func TestMyFunc(t *testing.T) {
	result := solutions.MyFunc()
	assert.Equal(t, "hello", result)
}
