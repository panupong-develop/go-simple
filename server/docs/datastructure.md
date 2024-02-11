# Hash Table
* Hash to find where to find the Node
* Node is a linked address that chains the Key-Value pair
  * key
  * value
  * *next address
* limit the number of crawling through the bucket
* The computation time won't grow with N. It's ~O(1)
* Think about knowing to find the book in the library that we already know the shelf number, 
  so the number of books on the shelf is limited.

A linked list Node
```
type Node struct {
    key   int
    value string
    next  *Node
}
```

Hash Table
```
type HashTable struct {
    capacity        int
    linkedListTable []*Node     // store the address
}
```

Constructor
```
func NewHashTable(capacity int) *HashTable {
    length := capacity
    return &HashTable{
        capacity: capacity,
        // a bucket slice of *Node
        linkedListTable: make([]*Node, length),
    }
}
```

Interface: Insert
```
func (ht *HashTable) Insert(key int, value string) {
    // find where to put the data
    index := key % ht.capacity

    // define the linked address
    node := &Node{key: key, value: value}

    if ht.table[index] == nil {
        ht.table[index] = node
    } else {
        current := ht.table[index]
        for current.next != nil {
            current = current.next  // craw the bucket
        }
        current.next = node
    }
}
```

Interface: Get
```
func (ht *HashTable) Get(key int) (string, bool) {
    index := key % capacity
    current := ht.table[index]

    for current != nil {
        if current.key == key {  // craw the bucket
            return current.value, true
        }
        current = current.next
    }
    return "", false
}
```
