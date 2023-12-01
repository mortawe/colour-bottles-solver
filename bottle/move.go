package bottle

func CheckMove(from, to Bottle) bool {
	if from.IsEmpty() || to.RemainingCapacity() < 1 {
		return false
	}
	if to.IsEmpty() || from.Peek() == to.Peek() {
		return true
	}
	return false
}

func Move(from, to *Bottle) bool {
	c, ok := from.Pop()
	if !ok {
		return false
	}
	return to.Push(c)
}
