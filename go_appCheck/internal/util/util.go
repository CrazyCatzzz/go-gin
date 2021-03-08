package util

import (
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

const timeFormat = "2006-01-02 15:04:05"

func TimeFormat(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(timeFormat)
}

func Copy(src, dest interface{}) {
	_ = copier.Copy(dest, src)
}

func GenUUID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

func String2Time(s string) (time.Time, error) {
	return time.ParseInLocation(timeFormat, s, time.Local)
}

func String2Int(s string) (int, error) {
	return strconv.Atoi(s)
}

func GetRoomByMac(s string) string {
	return strings.ToLower(strings.Join(strings.Split(s, ":"), ""))
}
