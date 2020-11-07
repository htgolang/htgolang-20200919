package main

type Less interface {
	Less(v interface{}) bool
}

type Equals interface {
	Equal(v interface{}) bool
}

type MapNode struct {
	Key   Equals
	Value Less
}

type Map []MapNode

func (m Map) Put(key Equals, value Less) {
	m = append(m, MapNode{key, value})
}

func (m Map) Get(key Equals) Less {
	for _, node := range m {
		if key.Equal(node.Key) {
			return node.Value
		}
	}
	return nil
}

func (m Map) Len() int {
	return len(m)
}

func (m Map) Less(i, j int) bool {
	return m[i].Value.Less(users[j].Value)
}

func (m Map) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
