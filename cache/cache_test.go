package cache

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testCache(t *testing.T, c Cache) {
	assert.Equal(t, nil, c.Get("1"))
	assert.Equal(t, nil, c.Get("2"))
	assert.Equal(t, nil, c.Get("3"))
	assert.Equal(t, 0, c.Len())

	c.Set("1", "foo")
	assert.Equal(t, "foo", c.Get("1"))
	assert.Equal(t, nil, c.Get("2"))
	assert.Equal(t, nil, c.Get("3"))
	assert.Equal(t, 1, c.Len())

	c.Set("3", "baz")
	assert.Equal(t, "foo", c.Get("1"))
	assert.Equal(t, nil, c.Get("2"))
	assert.Equal(t, "baz", c.Get("3"))
	assert.Equal(t, 2, c.Len())

	c.Set("2", "bar")
	assert.Equal(t, "foo", c.Get("1"))
	assert.Equal(t, "bar", c.Get("2"))
	assert.Equal(t, "baz", c.Get("3"))
	assert.Equal(t, 3, c.Len())

	c.Set("1", "quux")
	assert.Equal(t, "quux", c.Get("1"))
	assert.Equal(t, "bar", c.Get("2"))
	assert.Equal(t, "baz", c.Get("3"))
	assert.Equal(t, 3, c.Len())
}

func TestCaches(t *testing.T) {
	t.Run("Slice", func(t *testing.T) {
		testCache(t, NewSlice(0))
	})

	t.Run("SortedSlice", func(t *testing.T) {
		testCache(t, NewSortedSlice(0))
	})

	t.Run("Map", func(t *testing.T) {
		testCache(t, NewMap(0))
	})

	t.Run("SyncMap", func(t *testing.T) {
		testCache(t, NewSyncMap())
	})
}

var Sink interface{}

func benchmarkCache(b *testing.B, c Cache, items int) {
	ids := make([]string, items)
	for i, n := range rand.New(rand.NewSource(1)).Perm(items) {
		id := strconv.Itoa(n)
		c.Set(id, id)
		ids[i] = id
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, id := range ids {
			Sink = c.Get(id)
		}
	}
}

func BenchmarkCaches(b *testing.B) {
	const (
		minItems  = 0
		maxItems  = 20
		itemsStep = 2
		capHint   = 0
	)

	for items := minItems; items <= maxItems; items += itemsStep {
		b.Run(fmt.Sprintf("Slice,%d", items), func(b *testing.B) {
			benchmarkCache(b, NewSlice(capHint), items)
		})
	}

	for items := minItems; items <= maxItems; items += itemsStep {
		b.Run(fmt.Sprintf("SortedSlice,%d", items), func(b *testing.B) {
			benchmarkCache(b, NewSortedSlice(capHint), items)
		})
	}

	for items := minItems; items <= maxItems; items += itemsStep {
		b.Run(fmt.Sprintf("Map,%d", items), func(b *testing.B) {
			benchmarkCache(b, NewMap(capHint), items)
		})
	}

	for items := minItems; items <= maxItems; items += itemsStep {
		b.Run(fmt.Sprintf("SyncMap,%d", items), func(b *testing.B) {
			benchmarkCache(b, NewSyncMap(), items)
		})
	}
}
