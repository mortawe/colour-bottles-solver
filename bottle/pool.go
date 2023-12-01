package bottle

import (
	"sort"
	"strings"
)

type BottlePool []Bottle

func (bp BottlePool) Copy() BottlePool {
	bC := make([]Bottle, 0, len(bp))

	for _, e := range bp {
		bC = append(bC, e.copy())
	}
	return bC
}

func (bp BottlePool) String() string {
	var es []string
	for _, b := range bp {
		es = append(es, b.String())
	}
	return strings.Join(es, ",")
}

func (bp BottlePool) StringSorted() string {
	var es []string
	for _, b := range bp {
		es = append(es, b.String())
	}
	sort.Strings(es)
	return strings.Join(es, ",")
}

func (bp BottlePool) Check() bool {
	for _, b := range bp {
		if b.IsEmpty() {
			continue
		}
		if ok := b.Filled(); !ok {
			return false
		}
	}
	return true
}
