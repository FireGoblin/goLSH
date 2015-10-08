package main

type BucketIndex struct {
	word      string
	location  int
	suffixLen int
}

func (b *BucketIndex) largerNeighbors() []BucketIndex {
	if b.location == 2 {
		return nil
	}

	return []BucketIndex{BucketIndex{b.word, b.location, b.suffixLen + 1}, BucketIndex{b.word, b.location + 1, b.suffixLen}}
}
