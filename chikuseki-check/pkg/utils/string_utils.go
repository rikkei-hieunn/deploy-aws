package utils

import "chikuseki-check/model"

//ToString convert interface to string
func ToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	default:
		return model.EmptyString
	}
}
