package clock

import (
	"strings"
 	"strconv"
 	"time"
)

type Clock struct {
	hour, minute int
}

func New(hour, minute int) Clock {
	t := time.Date(0, time.November, 0, hour, minute, 0, 0, time.UTC)
	return Clock{hour: t.Hour(), minute: t.Minute()}
}

func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}

func (c Clock) Subtract(minutes int) Clock {
	return New(c.hour, c.minute-minutes)
}

func (c Clock) String() string {
	var b strings.Builder

	if c.hour < 10 {
		b.WriteString("0")
	}

	b.WriteString(strconv.FormatInt(int64(c.hour), 10))

	b.WriteString(":")

	if c.minute < 10 {
		b.WriteString("0")
	}

	b.WriteString(strconv.FormatInt(int64(c.minute), 10))
	return b.String()

}
