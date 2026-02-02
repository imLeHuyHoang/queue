package queue

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()
	if q == nil {
		t.Fatal("NewQueue returned nil")
	}
	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}
	if q.Len() != 0 {
		t.Errorf("New queue length should be 0, got %d", q.Len())
	}
}

func TestEnqueue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	if q.Len() != 3 {
		t.Errorf("Expected length 3, got %d", q.Len())
	}
	if q.IsEmpty() {
		t.Error("Queue should not be empty")
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	// Test FIFO order
	v, ok := q.Dequeue()
	if !ok || v != 10 {
		t.Errorf("Expected 10, got %d, ok=%v", v, ok)
	}

	v, ok = q.Dequeue()
	if !ok || v != 20 {
		t.Errorf("Expected 20, got %d, ok=%v", v, ok)
	}

	v, ok = q.Dequeue()
	if !ok || v != 30 {
		t.Errorf("Expected 30, got %d, ok=%v", v, ok)
	}

	// Test dequeue on empty queue
	v, ok = q.Dequeue()
	if ok {
		t.Error("Dequeue on empty queue should return ok=false")
	}
}

func TestFront(t *testing.T) {
	q := NewQueue[string]()

	// Test on empty queue
	_, ok := q.Front()
	if ok {
		t.Error("Front on empty queue should return ok=false")
	}

	q.Enqueue("first")
	q.Enqueue("second")

	// Test Front doesn't remove element
	v, ok := q.Front()
	if !ok || v != "first" {
		t.Errorf("Expected 'first', got %s, ok=%v", v, ok)
	}

	// Verify element is still there
	if q.Len() != 2 {
		t.Errorf("Front should not remove element, length=%d", q.Len())
	}

	v, ok = q.Front()
	if !ok || v != "first" {
		t.Errorf("Expected 'first' again, got %s, ok=%v", v, ok)
	}
}

func TestRear(t *testing.T) {
	q := NewQueue[int]()

	// Test on empty queue
	_, ok := q.Rear()
	if ok {
		t.Error("Rear on empty queue should return ok=false")
	}

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	v, ok := q.Rear()
	if !ok || v != 30 {
		t.Errorf("Expected 30, got %d, ok=%v", v, ok)
	}

	// Verify element is still there
	if q.Len() != 3 {
		t.Errorf("Rear should not remove element, length=%d", q.Len())
	}
}

func TestIsEmpty(t *testing.T) {
	q := NewQueue[int]()
	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}

	q.Enqueue(1)
	if q.IsEmpty() {
		t.Error("Queue with element should not be empty")
	}

	q.Dequeue()
	if !q.IsEmpty() {
		t.Error("Queue should be empty after removing all elements")
	}
}

func TestClear(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	q.Clear()

	if !q.IsEmpty() {
		t.Error("Queue should be empty after Clear")
	}
	if q.Len() != 0 {
		t.Errorf("Queue length should be 0 after Clear, got %d", q.Len())
	}
}

func TestToSlice(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	slice := q.ToSlice()

	if len(slice) != 3 {
		t.Errorf("Expected slice length 3, got %d", len(slice))
	}

	expected := []int{10, 20, 30}
	for i, v := range slice {
		if v != expected[i] {
			t.Errorf("At index %d: expected %d, got %d", i, expected[i], v)
		}
	}

	// Verify modifying slice doesn't affect queue
	slice[0] = 999
	front, _ := q.Front()
	if front != 10 {
		t.Error("Modifying returned slice should not affect queue")
	}
}

func TestQueueWithDifferentTypes(t *testing.T) {
	// Test with strings
	qs := NewQueue[string]()
	qs.Enqueue("hello")
	qs.Enqueue("world")

	v, ok := qs.Dequeue()
	if !ok || v != "hello" {
		t.Errorf("String queue: expected 'hello', got %s", v)
	}

	// Test with floats
	qf := NewQueue[float64]()
	qf.Enqueue(3.14)
	qf.Enqueue(2.71)

	vf, ok := qf.Front()
	if !ok || vf != 3.14 {
		t.Errorf("Float queue: expected 3.14, got %f", vf)
	}
}

func BenchmarkEnqueue(b *testing.B) {
	q := NewQueue[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

func BenchmarkDequeue(b *testing.B) {
	q := NewQueue[int]()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}
