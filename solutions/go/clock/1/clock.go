package clock

import (
    "fmt"
)

const minutesInDay = 1440

type Clock struct {
	minTotal int
}

func New(h, m int) Clock {
	total := extractMinutesTotal(h, m)
	return Clock{minTotal: total}
}

func (c Clock) Add(m int) Clock {
	totalAtualizado := extractMinutesTotal(0, c.minTotal+m)
	return Clock{minTotal: totalAtualizado}
}

func (c Clock) Subtract(m int) Clock {
	totalAtualizado := extractMinutesTotal(0, c.minTotal-m)
	return Clock{minTotal: totalAtualizado}
}

func (c Clock) String() string {
	h := c.minTotal / 60
	m := c.minTotal % 60
	return fmt.Sprintf("%02d:%02d", h, m)
}    

func extractMinutesTotal(h, m int) int {
	total := (h * 60) + m
	total = (total % minutesInDay + minutesInDay) % minutesInDay
	return total
}