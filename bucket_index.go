package main

type bucketIndex struct {
	word      string
	location  int
	suffixLen int
}

func (b *bucketIndex) largerNeighbor() bucketIndex {
	return bucketIndex{b.word, b.location, b.suffixLen + 1}
}
