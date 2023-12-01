package bottle

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/mortawe/colour-bottles-solver/stack"
)

type Bottle struct {
	s        stack.Stack[Colour]
	capacity int
}

func (b *Bottle) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Colours  []Colour
		Capacity int
	}{
		Colours:  b.s.Get(),
		Capacity: b.capacity,
	})
}

func (b *Bottle) UnmarshalJSON(bytes []byte) error {
	data := struct {
		Colours  []Colour
		Capacity int
	}{}
	err := json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}
	*b = New(data.Colours, data.Capacity)
	return nil
}

func (b *Bottle) String() string {
	var es []string
	for _, c := range b.s.Get() {
		es = append(es, strconv.Itoa(int(c)))
	}
	return strings.Join(es, "|")
}

func New(colours []Colour, capacity int) Bottle {
	b := Bottle{
		capacity: capacity,
	}
	for _, c := range colours {
		b.s.Push(c)
	}
	return b
}

// Peek returns top colour in bottle or `Empty` if bottle is empty.
func (b *Bottle) Peek() Colour {
	v, ok := b.s.Peek()
	if !ok {
		return Empty
	}
	return v
}

func (b *Bottle) Pop() (Colour, bool) {
	return b.s.Pop()
}

func (b *Bottle) Push(c Colour) bool {
	if b.s.Len() >= b.capacity {
		return false
	}
	top, ok := b.s.Peek()
	if ok && top != c {
		return false
	}
	b.s.Push(c)
	return true
}

func (b *Bottle) RemainingCapacity() int {
	return b.capacity - b.s.Len()
}

func (b *Bottle) IsEmpty() bool {
	return b.s.IsEmpty()
}

func (b *Bottle) GetTopCombination() (Colour, int) {
	colours := b.s.Get()
	if len(colours) == 0 {
		return Empty, 0
	}
	topColour := colours[0]
	var count int
	for _, e := range colours {
		if e == topColour {
			count++
		} else {
			break
		}
	}
	return topColour, count
}

func (b *Bottle) Filled() bool {
	_, count := b.GetTopCombination()
	return count == b.capacity
}

func ParseFile(filename string) (BottlePool, error) {
	byt, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var bottlePool BottlePool
	err = json.Unmarshal(byt, &bottlePool)

	if err != nil {
		return nil, err
	}
	return bottlePool, nil
}

func (b *Bottle) copy() Bottle {
	bC := Bottle{
		s:        b.s.Copy(),
		capacity: b.capacity,
	}
	return bC
}
