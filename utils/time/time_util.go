package time_util

import "time"
//格式化yyyy-MM-dd HH:mm:ss
func Format(time time.Time) string  {
	return time.Format("2006-01-02 03:04:05")
}
