package main

import "time"

type jsonTime struct {
	time.Time
	f string
}

func (j jsonTime) format() string {
	return j.Time.Format(j.f)
}

func (j jsonTime) MarshalText() ([]byte, error) {
	return []byte(j.format()), nil
}

func (j jsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + j.format() + `"`), nil
}
