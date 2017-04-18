package rcache

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
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
	t.Run("Map", func(t *testing.T) {
		testCache(t, NewMap(0))
	})
}

func benchmarkCache(b *testing.B, c Cache, items int) {
	ids := make([]string, items)
	for i, n := range rand.New(rand.NewSource(1)).Perm(items) {
		id := strconv.Itoa(n)
		ids[i] = id
	}

	var wg sync.WaitGroup

	{
		set := make(chan string, len(ids))
		for i := 0; i < runtime.GOMAXPROCS(0); i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for id := range set {
					c.Set(id, id)
				}
			}()
		}
		for _, id := range ids {
			set <- id
		}
		close(set)
		wg.Wait()
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		sink := make(chan interface{}, runtime.GOMAXPROCS(0))

		get := make(chan string, len(ids))
		for i := 0; i < runtime.GOMAXPROCS(0); i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				var s interface{}
				for id := range get {
					s = c.Get(id)
				}
				sink <- s
			}()
		}
		for _, id := range ids {
			get <- id
		}
		close(get)
		wg.Wait()
	}
}

func BenchmarkCaches(b *testing.B) {
	const maxItems = 10000

	b.Run(fmt.Sprintf("Map,%d", maxItems), func(b *testing.B) {
		benchmarkCache(b, NewMap(maxItems), maxItems)
	})
}
