package queue

import (
	"testing"
)

var s ItemQueue

func initQueue() *ItemQueue {
	if s.items == nil {
		s = ItemQueue{}
		s.New()
	}
	return &s
}

func TestEnqueue(t *testing.T) {
	s := initQueue()
	for i := 1; i <= 3; i++ {
		s.Enqueue(i)
	}

	if size := s.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
}

func TestDequeue(t *testing.T) {
	s.Dequeue()
	if size := len(s.items); size != 2 {
		t.Errorf("wrong count, exptect 2 and got %d", size)
	}

	s.Dequeue()
	s.Dequeue()
	if size := len(s.items); size != 0 {
		t.Errorf("wrong count, expected 0 and got %d", size)
	}

	if !s.IsEmpty() {
		t.Errorf("IsEmpty should return true")
	}
}