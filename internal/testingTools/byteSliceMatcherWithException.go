package testingTools

import (
	"go.uber.org/mock/gomock"
	"slices"
)

type ByteSliceMatcherWithException struct {
	values          []byte
	exceptedIndexes []int
}

func (m ByteSliceMatcherWithException) Matches(arg interface{}) bool {
	sarg := arg.([]byte)

	if len(sarg) != len(m.values) {
		return false
	}

	for index, b := range m.values {
		if !slices.Contains(m.exceptedIndexes, index) && b != sarg[index] {
			return false
		}
	}

	return true
}

func (m ByteSliceMatcherWithException) String() string {
	return "Should be " + string(m.values)
}

func NewByteSliceMatcherWithException(val []byte, ex []int) gomock.Matcher {
	return ByteSliceMatcherWithException{values: val, exceptedIndexes: ex}
}
