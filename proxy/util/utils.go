package util

import "strconv"

func ToLog(len int64) string {
	return strconv.FormatInt(len, 10)
}
