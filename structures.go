package main

import (
	"errors"
	"slices"
)

// Stack Implementation

type Stack struct {
	items []string
}

func (s *Stack) Push(items ...string) {
	slices.Reverse(items)
	s.items = append(s.items, items...)
}

func (s *Stack) Pop() (string, error) {
	if len(s.items) == 0 {
		return "", errors.New("Stack is empty")
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Set Implementation

type Set map[string]struct{}

func (s Set) Add(item string) {
	s[item] = struct{}{}
}

func (s Set) IsIn(item string) bool {
	_, exists := s[item]
	return exists
}

func (s Set) Delete(item string) {
	delete(s, item)
}

func (s Set) ToList() []string {
	var slice []string
	for key := range s {
		slice = append(slice, key)
	}
	return slice
}

// Queue Implementation

type Queue struct {
	elements []string
}

func (q *Queue) Add(elements ...string) {
	q.elements = append(q.elements, elements...)
}

func (q *Queue) Pop() string {
	element := q.elements[0]
	q.elements = q.elements[1:]
	return element
}

func (q *Queue) Peek() string {
	element := q.elements[0]
	return element
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue) Size() int {
	return len(q.elements)
}
