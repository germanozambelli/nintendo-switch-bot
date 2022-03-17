package action

type Combine struct {
	actions []Action
}

func NewCombine(actions []Action) Combine {
	return Combine{actions: actions}
}

func (combine Combine) Play() {
	for _, action := range combine.actions {
		action.Play()
	}
}
