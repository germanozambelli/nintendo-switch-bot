package action

import (
	"fmt"
	"time"
)

type Nothing struct {
	duration time.Duration
}

func NewNothing(duration time.Duration) Nothing {
	return Nothing{duration: duration}
}

func (nothing Nothing) Play() {
	fmt.Println(fmt.Sprintf("nothing for %d", nothing.duration))
	time.Sleep(nothing.duration)
}
