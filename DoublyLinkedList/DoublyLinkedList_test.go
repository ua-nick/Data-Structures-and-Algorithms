package DoublyLinkedList

import "testing"

func TestLinkedList_InsertFirst(t *testing.T) {
	ll := LinkedList{}

	var f, l int

	for i := 1; i <= 1000; i++ {
		ll.InsertFirst(i)
		if f, _ = ll.GetFirst(); f != i {
			t.Fatalf("GetFirst Error want:%d but return:%d", i, f)
		}
		if l, _ = ll.GetLast(); l != 1 {
			t.Fatalf("GetFirst Error want:%d but return:%d", 1, l)
		}
	}
}

func TestLinkedList_InsertLast(t *testing.T) {
	ll := LinkedList{}

	var f, l int

	for i := 1; i <= 1000; i++ {
		ll.InsertLast(i)
		if f, _ = ll.GetLast(); f != i {
			t.Fatalf("GetFirst Error want:%d but return:%d", i, f)
		}
		if l, _ = ll.GetFirst(); l != 1 {
			t.Fatalf("GetFirst Error want:%d but return:%d", 1, l)
		}
	}
}

func TestLinkedList_GetSize(t *testing.T) {
	ll := LinkedList{}
	count := 1000
	for i := 1; i <= count; i++ {
		ll.InsertFirst(i)
		if ll.GetSize() != i {
			t.Fatal()
		}
	}
}

func TestLinkedList_GetSize2(t *testing.T) {
	ll := LinkedList{}
	count := 1000
	for i := 1; i <= count; i++ {
		ll.InsertLast(i)
		if ll.GetSize() != i {
			t.Fatal()
		}
	}
}

func TestLinkedList_GetFirst(t *testing.T) {
	ll := LinkedList{}
	count := 1000
	for i := 1; i <= count; i++ {
		ll.InsertFirst(i)
	}

	if v, _ := ll.GetFirst(); v != count {
		t.Fatal()
	}
}

func TestLinkedList_GetLast(t *testing.T) {
	ll := LinkedList{}
	count := 1000
	for i := 1; i <= count; i++ {
		ll.InsertLast(i)
	}

	if v, _ := ll.GetLast(); v != count {
		t.Fatal()
	}
}

func TestLinkedList_GetItemsFromStart(t *testing.T) {
	ll := LinkedList{}
	count := 1000
	for i := 1; i <= count; i++ {
		ll.InsertFirst(i)
	}

	//from 1000 to 1
	for _, v := range ll.GetItemsFromStart() {
		if v != count {
			t.Fatal()
		}
		count--
	}
}

func TestLinkedList_GetItemsFromEnd(t *testing.T) {
	ll := LinkedList{}
	count := 1000
	for i := 1; i <= count; i++ {
		ll.InsertFirst(i)
	}

	//from 1 to 1000
	index := 1
	for _, v := range ll.GetItemsFromEnd() {
		if v != index {
			t.Fatal()
		}
		index++
	}
}

func TestLinkedList_SearchValue(t *testing.T) {
	ll := LinkedList{}
	count := 1000
	for i := 1; i <= count; i++ {
		ll.InsertFirst(i)
	}

	for i := 1; i <= count; i++ {
		if !ll.SearchValue(i) {
			t.Fatal()
		}
	}

	if ll.SearchValue(1001) {
		t.Fatal()
	}
	if ll.SearchValue(0) {
		t.Fatal()
	}
}

func TestLinkedList_RemoveByValue(t *testing.T) {
	ll := LinkedList{}
	count := 1000
	for i := 1; i <= count; i++ {
		ll.InsertFirst(i)
	}

	for i := 1; i <= count; i++ {
		ll.RemoveByValue(i)
		if ll.SearchValue(i) {
			t.Fatal()
		}
	}
}

func TestLinkedList_RemoveByValue1(t *testing.T) {
	ll := LinkedList{}
	count := 1000
	for i := 1; i <= count; i++ {
		ll.InsertLast(i)
	}

	for i := 1; i <= count; i++ {
		ll.RemoveByValue(i)
		if ll.SearchValue(i) {
			t.Fatal()
		}
	}
}

func TestLinkedList_RemoveByIndex(t *testing.T) {
	ll := LinkedList{}
	count := 1000
	for i := 1; i <= count; i++ {
		ll.InsertLast(i)
	}

	//from 1 to 1000
	for i := 1; i <= count; i++ {
		ll.RemoveByIndex(0)
		if ll.SearchValue(i) {
			t.Fatal()
		}
	}
}

func TestLinkedList_RemoveByIndex2(t *testing.T) {
	ll := LinkedList{}
	count := 1000
	for i := 1; i <= count; i++ {
		ll.InsertFirst(i)
	}

	//from 1 to 1000
	for i := 1; i <= count; i++ {
		ll.RemoveByIndex(ll.GetSize() - 1)
		if ll.SearchValue(i) {
			t.Fatal()
		}
	}
}
