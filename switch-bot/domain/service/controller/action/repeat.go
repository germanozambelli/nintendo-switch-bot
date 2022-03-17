package action

type Repeat struct {
	action Action
	times  int
}

func NewRepeat(action Action, times int) Repeat {
	return Repeat{action: action, times: times}
}

func (repeat Repeat) Play() {
	for i := 0; i < repeat.times; i++ {
		repeat.action.Play()
	}
}
