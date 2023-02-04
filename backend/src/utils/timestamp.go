package utils

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

type Timestamp struct {
	Time  time.Time
	Valid bool
}

func (ts *Timestamp) UnmarshalJSON(data []byte) error {
	newTime, err := time.Parse(timeFormat, strings.Trim(string(data), "\""))
	if err != nil {
		return err
	}

	*ts = Timestamp{Time: newTime, Valid: true}
	return nil
}

func (ts *Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", time.Time((*ts).Time).Format(timeFormat))), nil
}

func (ts Timestamp) Value() (driver.Value, error) {
	return time.Time(ts.Time), nil
}

func (ts *Timestamp) Scan(src interface{}) error {
	if val, ok := src.(time.Time); ok {
		*ts = Timestamp{Time: val, Valid: true}
	} else {
		*ts = Timestamp{Time: time.Now(), Valid: false}
	}

	return nil
}
