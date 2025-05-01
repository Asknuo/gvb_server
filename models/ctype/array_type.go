package ctype

import (
	"errors"
	"strings"
)

type Array []string

func (t *Array) Scan(value interface{}) error {
	if value == nil {
		*t = []string{}
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	default:
		return errors.New("unsupported type for ctype.Array, expected string or []byte")
	}

	if str == "" {
		*t = []string{}
		return nil
	}
	*t = strings.Split(str, "\n")
	return nil
}
