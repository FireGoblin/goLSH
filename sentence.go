package main

type Sentence []strings

func (s Sentence) buckets() []BucketIndex {
	if s[0] == s[1] {
		return []BucketIndex{s[0], 0, len(s) - 1}
	}
	return []BucketIndex{{s[0], 0, len(s) - 1}, {s[1], 1, len(s) - 2}}
}

func (s Sentence) compareWithSameLength(target Sentence) bool {
	mismatchAvailable = true
	for i, _ := range s {
		if s[i] != target[i] {
			if !mismatchAvailable {
				return false
			}
			mismatchAvailable = false
		}
	}
	return true
}

func (s Sentence) compareWithLonger(target Sentence) bool {
	offset := 0
	for i, _ := range s {
		if s[i] != target[i+offset] {
			if offset == 1 {
				return false
			}
			offset++
		}
	}
	return true
}
