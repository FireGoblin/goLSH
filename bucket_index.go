package main

type BucketIndex struct {
	word      string
	location  int
	suffixLen int
}

func (b *BucketIndex) largerNeighbor() BucketIndex {
	return BucketIndex{b.word, b.location, b.suffixLen + 1}
}
