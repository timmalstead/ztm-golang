package queue

import "testing"

const QueueAmount = 3

var queueToTest Queue

func TestAddQueue(t *testing.T) {
	queueToTest = NewQueue(QueueAmount)
	for i := 0; i < QueueAmount; i++ {
		if len(queueToTest.items) != i {
			t.Errorf("Incorrect queue count: %v. Desired: %v", len(queueToTest.items), i)
		}
		if !queueToTest.AppendItem(i) {
			t.Errorf("Failed to append item %v to queue", i)
		}
	}
	// test after queue is full
	if queueToTest.AppendItem(QueueAmount) {
		t.Errorf("Adding item to queue beyond capacity!")
	}
}

func TestNextItem(t *testing.T) {
	queueToTest = NewQueue(QueueAmount)
	for i := 0; i < QueueAmount; i++ {
		queueToTest.AppendItem(i)
	}

	for i := 0; i < QueueAmount; i++ {
		var item, ok = queueToTest.NextItem()
		if !ok {
			t.Errorf("queue next item not returning correctly, investigate error")
		}
		if item != i {
			t.Errorf("item %v is not equal to expected %v", item, i)
		}
	}

	// queue should not empty
	var item, ok = queueToTest.NextItem()
	if ok {
		t.Errorf("queue should be empty")
	}
	if item != 0 {
		t.Errorf("empty queue should return 0, instead of %v", item)
	}

}
