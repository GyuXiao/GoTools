package cmd

import (
	"github.com/spf13/cobra"
	"goTools/internal/word"
	"log"
	"strings"
)

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamelCase
	ModeUnderscoreToLowerCamelCase
	ModeCamelCaseToUnderscore
)

var desc = strings.Join([]string{
	"子命令支持各种单词格式转换，模式如下所示：",
	"1：全部转大写",
	"2：全部转小写",
	"3：下划线转大写驼峰",
	"4：下划线转小写驼峰",
	"5：驼峰转下划线",
}, "\n")

// 单词格式转换子命令

var mode int
var str string

// 验证命令：go run main.go help words

var wordCmd = &cobra.Command{
	Use:   "words",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持此转换模式，请执行 help word 查看帮助")
		}
		log.Printf("输出结果：%s", content)
	},
}

func init() {
	// 验证命令：
	// go run main.go word -s=XiaoKai -m=2
	// go run main.go word -s=xiaokai -m=5

	// 根据单词转换所需要的参数，分别是单词内容和转换的模式进行命令行参数的设置和初始化
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入内容")
	wordCmd.Flags().IntVarP(&mode, "mode", "m", 0, "请输入转换模式")
}
