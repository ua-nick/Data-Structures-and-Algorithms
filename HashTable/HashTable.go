package HashTable

import "hash/fnv"

type TableItem struct {
	key  string
	data int
	next *TableItem
}

type HashTable struct {
	data [256]*TableItem
}

func (table *HashTable) Add(key string, i int) {
	position := generateHash(key)
	if table.data[position] == nil {
		table.data[position] = &TableItem{key: key, data: i}
		return
	}
	current := table.data[position]
	for current.next != nil {
		current = current.next
	}
	current.next = &TableItem{key: key, data: i}
}

func (table *HashTable) Get(key string) (int, bool) {
	position := generateHash(key)
	current := table.data[position]
	for current != nil {
		if current.key == key {
			return current.data, true
		}
		current = current.next
	}
	return 0, false
}

func (table *HashTable) Set(key string, value int) bool {
	position := generateHash(key)
	current := table.data[position]
	for current != nil {
		if current.key == key {
			current.data = value
			return true
		}
		current = current.next
	}
	return false
}

func (table *HashTable) Remove(key string) bool {
	position := generateHash(key)
	if table.data[position] == nil {
		return false
	}
	if table.data[position].key == key {
		table.data[position] = table.data[position].next
		return true
	}
	current := table.data[position]
	for current.next != nil {
		if current.next.key == key {
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

func generateHash(s string) uint8 {
	hash := fnv.New32a()
	hash.Write([]byte(s))
	return uint8(hash.Sum32() % 256)
}
