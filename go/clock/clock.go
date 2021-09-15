package clock

import "fmt"

type Clock struct {
	Hour   int
	Minute int
}

func New(hour, minute int) Clock {
	hour += minute / 60
	minute = minute % 60

	if minute < 0 {
		hour -= 1
		minute += 60
	}

	hour = hour % 24
	if hour < 0 {
		hour += 24
	}
	return Clock{hour, minute}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}

func (c Clock) Add(minutes int) Clock {
	c.Minute += minutes
	return New(c.Hour, c.Minute)
}

func (c Clock) Subtract(minutes int) Clock {
	c.Minute -= minutes
	return New(c.Hour, c.Minute)
}
