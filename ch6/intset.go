package ch6

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(x ...int) {
	for i := range x {
		s.Add(x[i])
	}
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word > len(s.words) {
		return
	}
	s.words[word] &= ^(1 << bit)
}

func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] &= 0
	}
}

func (s *IntSet) Copy() *IntSet {
	arr := make([]uint64, len(s.words))
	for i, word := range s.words {
		arr[i] = word
	}
	return &IntSet{
		words: arr,
	}
}

//UnionWith 并集
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//IntersectWith 交集
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
	if len(t.words) < len(s.words) {
		for i := len(t.words); i < len(s.words); i++ {
			s.words[i] &= 0
		}
	}
}

//DifferenceWith A中出现B中没出现的
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] = s.words[i] ^ tword
			s.words[i] &= ^tword
		}
	}
}

//SymmetricDifference A中出现B中没出现的，或A中没出现B中出现的
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] = s.words[i] ^ tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) Len() int {
	len := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if (word & (1 << j)) != 0 {
				len++
			}
		}
	}
	return len
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
