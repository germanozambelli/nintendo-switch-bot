package value_object

type State string

const (
	TALKING State = "TALKING"
	WAITING       = "WAITING"
	MOVING        = "MOVING"
	OTHER         = "OTHER"
)
