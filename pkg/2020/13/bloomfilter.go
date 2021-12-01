package day13

type BloomFilter struct {
	m       int64
	k       int
	hashFns []func(int64) []int8
}

func NewBloomFilter() *BloomFilter {
	return &BloomFilter{}
}

func (b *BloomFilter) Add(element int64) bool {
	return false
}

func (b *BloomFilter) Query(element int64) bool {
	return false
}
