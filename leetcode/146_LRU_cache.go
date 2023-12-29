package questions

type LRUCache struct {
	cache   map[int]*node
	freHead *node
	freTail *node
	cap     int
	curLen  int
}

type node struct {
	k    int
	v    int
	prev *node
	next *node
}

func Constructor(capacity int) LRUCache {
	instance := LRUCache{
		cache:   make(map[int]*node),
		freHead: &node{k: -1, v: -1},
		freTail: &node{k: -1, v: -1},
		cap:     capacity,
		curLen:  0,
	}
	instance.freHead.next = instance.freTail
	instance.freTail.prev = instance.freHead
	return instance
}

func (this *LRUCache) Get(key int) int {
	res := this.get(key)
	if res == nil {
		return -1
	}
	return res.v
}

func (this *LRUCache) get(key int) *node {
	if _, ok := this.cache[key]; !ok {
		return nil
	}
	target := this.cache[key]
	oldLeft := target.prev
	oldRight := target.next
	newLeft := this.freTail.prev
	newRight := this.freTail
	if newLeft != target {
		oldLeft.next = oldRight
		oldRight.prev = oldLeft
		newLeft.next = target
		target.prev = newLeft
		target.next = newRight
		newRight.prev = target
	}
	return target
}

func (this *LRUCache) Put(key int, value int) {
	var target *node
	if _, ok := this.cache[key]; ok {
		target = this.get(key)
		target.v = value
	} else {
		target = &node{k: key, v: value}
		if this.curLen < this.cap {
			// insert the new node
			newLeft := this.freTail.prev
			newRight := this.freTail
			newLeft.next = target
			target.prev = newLeft
			target.next = newRight
			newRight.prev = target
			this.curLen += 1
		} else {
			// delete the old node
			oldLeft := this.freHead
			oldTarget := oldLeft.next
			oldRight := oldTarget.next
			oldLeft.next = oldRight
			oldRight.prev = oldLeft
			// insert the new node
			newRight := this.freTail
			newLeft := newRight.prev
			newLeft.next = target
			target.prev = newLeft
			target.next = newRight
			newRight.prev = target
			// delete from cache
			delete(this.cache, oldTarget.k)
		}
	}
	this.cache[key] = target
}
