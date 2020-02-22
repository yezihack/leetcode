package linked

//定义节点
type Node struct {
	key, value int   //key:value
	prev, next *Node //前驱指针与后继指针
}

//定义LRU结构体
type LRUCache struct {
	capacity  int           //上限容量
	MapCache  map[int]*Node //map缓存节点,实现O(1)操作
	head, end *Node         //头节点与尾结点
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		MapCache: make(map[int]*Node), //初使化,map必须初使化才能用.
		head:     nil,
		end:      nil,
	}
}

//获取操作
//1. 判断key是否存在
//---存在:从缓存里获取节点信息,然后删除,再添加到链表的头部位置.
//---不存在:直接返回-1
func (this *LRUCache) Get(key int) int {
	node, ok := this.MapCache[key] //从缓冲中取出
	if !ok {                       //缓存不存在则返回-1
		return -1
	}
	this.remove(node)    //删除节点
	this.setHeader(node) //添加到头节点
	return node.value    //返回数据.
}
func (this *LRUCache) Put(key, value int) {
	if node, ok := this.MapCache[key]; ok {
		node.value = value
		this.remove(node)
		this.setHeader(node)
	} else {
		if len(this.MapCache) >= this.capacity {
			delete(this.MapCache, key) //删除缓存中的数据.
			this.remove(this.end)      //再移除最后一个节点
		}
		node := &Node{key: key, value: value}
		this.MapCache[key] = node
		this.setHeader(node)
	}
}

//移除节点
//1.判断前驱是否为空
//是:此节点就是头节点.
//否:不是头节点,则node.pre.next = node.next
//2.判断后继是否为空
//是:此节点就是尾节点.
//否:不是尾节点,则node.next.pre = node.pre
func (this *LRUCache) remove(node *Node) {
	if node.prev == nil {
		this.head = node.next
	} else {
		node.prev.next = node.next
	}
	if node.next == nil {
		this.end = node.prev
		this.end.next = nil
	} else {
		node.next.prev = node.prev
	}
}

//设置头节点.
//头节点的特性: 前驱指针为空.
//1.判断头节点head是否为空
//是: 直接设置
//否: 将当前的节点设置为第二个节点, node设置为新的头节点.
//2.判断end尾节点是否为空
//是:直接设置
//否:不操作
func (this *LRUCache) setHeader(node *Node) {
	node.prev = nil
	if this.head == nil {
		this.head = node
	} else {
		head := this.head
		node.next = head
		head.prev = node
		this.head = node
	}
	if this.end == nil {
		this.end = node
	}
}
