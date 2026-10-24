package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	hour int
	minute int
}

func New(h, m int) Clock {
	total := h * 60 + m
	// 24 hours in a day
	total = ((total % 1440) + 1440) % 1440
	return Clock {
		hour: total / 60,
		minute: total % 60,
	}
}

func (c Clock) Add(m int) Clock {
	return New(c.hour, c.minute+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hour, c.minute-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
