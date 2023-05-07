package timer

import "time"

func GtNowTime() time.Time {
	return time.Now()
}

// 返回 currentTimer + d（d 可以是正数或负数）

func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	// ParseDuration 从字符串中解析出 duration（持续时间）
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	// 求出 当前时间+duration 后所得到的最终时间
	return currentTimer.Add(duration), nil
}

