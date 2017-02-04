package nex

import (
	"strconv"
)

func (f *uniform) Int(key string) int {
	value := f.Get(key)
	if value == "" {
		return 0
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return v
}

func (f *uniform) IntOrDefault(key string, def int) int {
	value := f.Get(key)
	if value == "" {
		return 0
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		return def
	}
	return v
}

func (f *uniform) Int64(key string) int64 {
	value := f.Get(key)
	if value == "" {
		return 0
	}
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0
	}
	return v
}

func (f *uniform) Int64OrDefault(key string, def int64) int64 {
	value := f.Get(key)
	if value == "" {
		return 0
	}
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return def
	}
	return v
}

func (f *uniform) Uint64(key string) uint64 {
	value := f.Get(key)
	if value == "" {
		return 0
	}
	v, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0
	}
	return v
}

func (f *uniform) Uint64OrDefault(key string, def uint64) uint64 {
	value := f.Get(key)
	if value == "" {
		return 0
	}
	v, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return def
	}
	return v
}
