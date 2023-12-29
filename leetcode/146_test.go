package questions

import "testing"

func Test146(t *testing.T) {
	c := Constructor(2)
	c.Put(2, 1)
	c.Put(2, 2)
	c.Get(2)
	c.Put(1, 1)
	c.Put(4, 1)
	c.Get(2)
}
