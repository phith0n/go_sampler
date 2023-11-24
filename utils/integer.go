package utils

import (
	"fmt"
	"strconv"
)

func EnsureUInt(value interface{}) uint {
	var s string
	switch t := value.(type) {
	case string:
		s = t
	case int8, int16, int32, int64, uint8, uint16, uint32, uint64, int:
		s = fmt.Sprintf("%v", t)
	case uint:
		return t
	default:
		return 0
	}

	r, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		r = 0
	}

	if uint(r) < 0 {
		return 0
	} else {
		return uint(r)
	}
}
