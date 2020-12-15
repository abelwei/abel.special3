package general

import (
	"strconv"
	"time"
)

type Timelib struct {
	Current time.Time
}

func NewTime() *Timelib{
	return &Timelib{
		Current: time.Now(),
	}
}

func (self *Timelib) Timestamp() string {
	timestamp := strconv.FormatInt(self.Current.UTC().UnixNano(), 10)
	timestamp = timestamp[:10]
	return timestamp
}

func (self *Timelib) Format() string{
	return self.Current.Format("2006-01-02 15:04:05")
}
