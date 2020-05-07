package _146_lru_cache

// 首先是通过双向链表来存放数据，插入和删除实现O(1)。分别记录链表的header tail指针
// 对于查找链表是无法实现O(1)，需要借助于map来实现
// 为了在链表操作过程中更少的判断头尾节点是否为null，采用哨兵机制，头尾都添加哨兵.
// 这里只是实现一个线程不安全的版本

type LRUCache struct {
	cap    int
	header *Node
	tail   *Node
	m      map[int]*Node
}

type Node struct {
	key   int
	value int
	pre   *Node
	next  *Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		cap:    capacity,
		header: &Node{},
		tail:   &Node{},
		m:      make(map[int]*Node, capacity)}
	cache.header.next = cache.tail
	cache.tail.pre = cache.header
	return cache
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; !ok {
		return -1
	} else {
		// 存在的化需要把这个node移动到header.不需要关心是否是第一个。有哨兵机制
		this.remove(node)
		this.putHead(node)
		return node.value
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		//　存在节点.更新数据。移动到head
		node.value = value
		this.remove(node)
		this.putHead(node)
	} else {
		// 不存在
		// 如果容器已经满了，需要删除链表尾部,从map中删除
		if len(this.m) >= this.cap {
			deletekey := this.tail.pre.key
			this.remove(this.tail.pre)
			delete(this.m, deletekey)
		}

		//创建新的节点,并放在head,添加到map中
		newNode := Node{key: key, value: value}
		this.putHead(&newNode)
		this.m[key] = &newNode
	}
}

func (this *LRUCache) remove(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) putHead(node *Node) {
	originNext := this.header.next
	this.header.next = node
	node.next = originNext

	originNext.pre = node
	node.pre = this.header
}

// 大佬的版本
//// LRUCache contains a hash map and a doubly linked list
//type LRUCache struct {
//	cap int                   // capacity
//	l   *list.List            // doubly linked list
//	m   map[int]*list.Element // hash table for checking if list node exists
//}
//
//// Pair is the value of a list node.
//type Pair struct {
//	key   int
//	value int
//}
//
//// Constructor initializes a new LRUCache.
//func Constructor(capacity int) LRUCache {
//	return LRUCache{
//		cap: capacity,
//		l:   new(list.List),
//		m:   make(map[int]*list.Element, capacity),
//	}
//}
//
//// Get a list node from the hash map.
//func (c *LRUCache) Get(key int) int {
//	// check if list node exists
//	if node, ok := c.m[key]; ok {
//		val := node.Value.(*list.Element).Value.(Pair).value
//		// move node to front
//		c.l.MoveToFront(node)
//		return val
//	}
//	return -1
//}
//
//// Put key and value in the LRUCache
//func (c *LRUCache) Put(key int, value int) {
//	// check if list node exists
//	if node, ok := c.m[key]; ok {
//		// move the node to front
//		c.l.MoveToFront(node)
//		// update the value of a list node
//		node.Value.(*list.Element).Value = Pair{key: key, value: value}
//	} else {
//		// delete the last list node if the list is full
//		if c.l.Len() == c.cap {
//			// get the key that we want to delete
//			idx := c.l.Back().Value.(*list.Element).Value.(Pair).key
//			// delete the node pointer in the hash map by key
//			delete(c.m, idx)
//			// remove the last list node
//			c.l.Remove(c.l.Back())
//		}
//		// initialize a list node
//		node := &list.Element{
//			Value: Pair{
//				key:   key,
//				value: value,
//			},
//		}
//		// push the new list node into the list
//		ptr := c.l.PushFront(node)
//		// save the node pointer in the hash map
//		c.m[key] = ptr
//	}
//}