package utils

import (
	"fmt"
	"reflect"
)

type Set[K comparable] map[K]struct{}

func NewSet[T comparable](args ...T) Set[T] {
	set := make(Set[T])
	for _, arg := range args {
		set[arg] = struct{}{}
	}
	return set
}

func (s Set[K]) Has(args ...K) bool {
	for _, arg := range args {
		_, has := s[arg]
		if !has {
			return false
		}
	}
	return true
}

func (s *Set[K]) Add(args ...K) {
	for _, arg := range args {
		(*s)[arg] = struct{}{}
	}
}

func (s *Set[K]) Delete(args ...K) {
	for _, arg := range args {
		delete((*s), arg)
	}
}

func (S Set[K]) ToList() []K {
	l := len(S)
	outList := make([]K, l)
	var i int
	for s := range S {
		outList[i] = s
		i++
	}
	return outList
}

func (S Set[K]) String() string {
	l := S.ToList()
	s := fmt.Sprintf("%v", l)
	s = s[1 : len(s)-1]
	return fmt.Sprintf("{%s}", s)
}

func (S Set[K]) Equals(T Set[K]) bool {
	return reflect.DeepEqual(S, T)
}
