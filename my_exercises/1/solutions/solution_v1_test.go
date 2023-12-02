package solutions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/gustavooferreira/coding/my_exercises/1/solutions"
)

func TestParenthesisMatch(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := solutions.ParenthesisMatch(tc.str)
			assert.Equal(t, tc.expectedResult, result)
		})
	}

}

func TestRuneStack(t *testing.T) {
	t.Run("test pushing and poping from stack", func(t *testing.T) {
		// create stack
		stack := solutions.RunesStack{}

		assert.True(t, stack.Empty())

		stack.Push('a')
		stack.Push('b')
		stack.Push('c')

		assert.Equal(t, 3, stack.Length())

		r, err := stack.Pop()
		require.NoError(t, err)
		assert.Equal(t, 'c', r)
		r, err = stack.Pop()
		require.NoError(t, err)
		assert.Equal(t, 'b', r)
		r, err = stack.Pop()
		require.NoError(t, err)
		assert.Equal(t, 'a', r)

		assert.True(t, stack.Empty())

		_, err = stack.Pop()
		require.Error(t, err)
	})
}
