package sets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	s := NewSet(1, 2, 3, 3, 1, 2, 2)
	assert.Equal(t, 3, s.Len())
	assert.Equal(t, []int{1, 2, 3}, s.Elements())
}

func TestSet_Contains(t *testing.T) {
	s := NewSet(1, 2, 3, 3)
	assert.True(t, s.Contains(1))
	assert.False(t, s.Contains(4))
}

func TestSet_Union(t *testing.T) {
	s1 := NewSet(1, 2, 3)
	s2 := NewSet(3, 4, 5)
	s3 := s1.Union(s2)
	assert.Equal(t, 5, s3.Len())
	assert.Equal(t, []int{1, 2, 3, 4, 5}, s3.Elements())
}

func TestSet_Intersect(t *testing.T) {
	s1 := NewSet(1, 2, 3)
	s2 := NewSet(3, 4, 5)
	s3 := s1.Intersect(s2)
	assert.Equal(t, 1, s3.Len())
	assert.Equal(t, []int{3}, s3.Elements())
}

func TestSet_Difference(t *testing.T) {
	s1 := NewSet(1, 2, 3)
	s2 := NewSet(3, 4, 5)
	s3 := s1.Difference(s2)
	assert.Equal(t, 2, s3.Len())
	assert.Equal(t, []int{1, 2}, s3.Elements())
}
