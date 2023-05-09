package cmd

import (
	"github.com/spf13/cobra"
	"goTools/internal/timer"
	"log"
	"strconv"
	"strings"
	"time"
)

var calculateTime string
var duration string
var layout = "2006-01-02 15:04:05"

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GtNowTime()
		// 获取当前时间的 Time 对象后，一共输出两个不同格式的时间
		// 1，第一个格式：通过调用 Format 方法设定约定的 2006-01-02 15:04:05 格式来进行时间的标准格式化。
		// 2，第二个格式：通过调用 Unix 方法返回 Unix 时间，也就是时间戳
		log.Printf("输出结果：%s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		currentTimer := getCurrentTime()
		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}
		log.Printf("输出结果: %s, %d", t.Format(layout), t.Unix())
	},
}

func getCurrentTime() time.Time {
	var currentTimer time.Time
	// 如果待计算时间为空，则当前时间作为待计算时间
	if calculateTime == "" {
		currentTimer = timer.GtNowTime()
	} else {
		spaceCnt := strings.Count(calculateTime, " ")
		if spaceCnt == 0 {
			layout = "2006-01-02"
		}
		if spaceCnt == 1 {
			layout = "2006-01-02 15:04:05"
		}
		var err error
		currentTimer, err = time.Parse(layout, calculateTime)
		if err != nil {
			t, _ := strconv.Atoi(calculateTime)
			currentTimer = time.Unix(int64(t), 0)
		}
	}
	return currentTimer
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)
	// 验证命令
	// go run main.go time now
	// go run main.go time calc -c="2023-05-06 22:34:32" -d=5m
	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "需要计算的时间")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", ` 持续时间，有效时间单位为"ns", "us" (or "µ s"), "ms", "s", "m", "h"`)
}
