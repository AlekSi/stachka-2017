package rcache

type item struct {
	id    string
	value interface{}
}

type Cache interface {
	Get(id string) interface{}
	Set(id string, value interface{})
	Len() int
}
