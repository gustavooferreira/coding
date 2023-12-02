package ops_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gustavooferreira/coding/atomicops/ops"
)

// TestIncrementOperation is only here so we can play with the -bench and -run flags for go test.
func TestIncrementOperation(t *testing.T) {
	op := ops.Operation{}
	op.Increment()
	result := op.Get()
	t.Logf("the result is %d", result)
	assert.Equal(t, 1, result)
}

func TestIncrementOperationSubtests(t *testing.T) {
	t.Run("my test should return 2", func(t *testing.T) {
		op := ops.Operation{}
		op.Increment()
		op.Increment()
		result := op.Get()
		t.Logf("the result is %d", result)
		assert.Equal(t, 2, result)
	})

	t.Run("my test should return 0", func(t *testing.T) {
		op := ops.Operation{}
		result := op.Get()
		t.Logf("the result is %d", result)
		assert.Equal(t, 0, result)
	})
}

func BenchmarkIncrement(b *testing.B) {
	op := ops.Operation{}
	for i := 0; i < b.N; i++ {
		op.Increment()
	}
}

func BenchmarkIncrementAtomicInteger(b *testing.B) {
	op := ops.OperationAtomicInteger{}
	for i := 0; i < b.N; i++ {
		op.Increment()
	}
}

func BenchmarkIncrementMutex(b *testing.B) {
	op := ops.OperationMutex{}
	for i := 0; i < b.N; i++ {
		op.Increment()
	}
}

func BenchmarkIncrementAtomicValue(b *testing.B) {
	op := ops.OperationAtomicValue{}
	for i := 0; i < b.N; i++ {
		op.Increment()
	}
}

func BenchmarkGet(b *testing.B) {
	op := ops.Operation{}
	for i := 0; i < b.N; i++ {
		op.Get()
	}
}

func BenchmarkGetAtomicInteger(b *testing.B) {
	op := ops.OperationAtomicInteger{}
	for i := 0; i < b.N; i++ {
		op.Get()
	}
}

func BenchmarkGetMutex(b *testing.B) {
	op := ops.OperationMutex{}
	for i := 0; i < b.N; i++ {
		op.Get()
	}
}

func BenchmarkGetAtomicValue(b *testing.B) {
	op := ops.OperationAtomicValue{}
	// set value so that it doesnt return nil
	op.Increment()
	for i := 0; i < b.N; i++ {
		op.Get()
	}
}
